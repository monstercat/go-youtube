package youtube

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/google"
)

// NewClient creates an authenticated *http.Client from a Google service account
// JWT key (JSON). The returned client automatically handles token refresh and
// attaches the Bearer token to every outgoing request.
//
// keyJSON is the raw bytes of the service account key file (the JSON downloaded
// from the Google Cloud Console). scopes specifies the OAuth2 scopes the client
// should request — for the Content ID API this is typically:
//
//   - https://www.googleapis.com/auth/youtubepartner
//
// For the YouTube Data API:
//
//   - https://www.googleapis.com/auth/youtube
//   - https://www.googleapis.com/auth/youtube.readonly
func NewClient(keyJSON []byte, scopes []string) (*http.Client, error) {
	cfg, err := google.JWTConfigFromJSON(keyJSON, scopes...)
	if err != nil {
		return nil, err
	}
	return cfg.Client(context.Background()), nil
}
