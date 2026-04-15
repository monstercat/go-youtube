package youtube

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	// EnvYouTubeJWT is the environment variable that provides the JWT
	// credentials. It can be either a file path to a JSON key file or the
	// raw JSON string itself.
	EnvYouTubeJWT = "YOUTUBE_JWT"
)

// readJWT reads the JWT credentials from the environment variable. If the
// value is a path to an existing file, the file contents are returned.
// Otherwise the value is treated as the raw JSON key string.
func readJWT(t *testing.T) []byte {
	t.Helper()

	raw := os.Getenv(EnvYouTubeJWT)
	require.NotEmpty(t, raw, "environment variable %s must be set", EnvYouTubeJWT)

	// A JSON key string starts with '{'; a file path never does.
	if len(raw) > 0 && raw[0] == '{' {
		return []byte(raw)
	}

	// Otherwise treat it as a file path.
	data, err := os.ReadFile(raw)
	require.NoError(t, err, "YOUTUBE_JWT does not start with '{' and could not be read as a file path: %s", raw)
	return data
}

// newTestRunner creates a CustomClientRunner authenticated with the JWT
// from the environment. It skips the test if YOUTUBE_JWT is not set.
func newTestRunner(t *testing.T) *CustomClientRunner {
	t.Helper()

	if os.Getenv(EnvYouTubeJWT) == "" {
		t.Skipf("skipping: %s not set", EnvYouTubeJWT)
	}

	keyJSON := readJWT(t)
	client, err := NewClient(keyJSON, []string{
		YoutubeForceSslScope,
		YoutubeScope,
		YoutubepartnerScope,
	})
	require.NoError(t, err, "failed to create authenticated client")

	return &CustomClientRunner{Client: client}
}

func TestNewClient(t *testing.T) {
	runner := newTestRunner(t)
	require.NotNil(t, runner.Client)
}
