package youtube

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func errResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

// TestDecodeResponse_ContentIDEnvelope verifies the Google API error
// envelope (object-valued "error" with an errors[] array) is parsed into
// the structured Reason / Errors fields.
func TestDecodeResponse_ContentIDEnvelope(t *testing.T) {
	body := `{"error":{"code":400,"message":"The claimed asset is not owned by the caller.","errors":[{"message":"The claimed asset is not owned by the caller.","domain":"youtubePartner.claims.update","reason":"missingOwnership"}]}}`

	err := DecodeResponse(errResponse(http.StatusBadRequest, body), nil)
	require.Error(t, err)

	var ytErr Error
	require.ErrorAs(t, err, &ytErr)
	require.Equal(t, http.StatusBadRequest, ytErr.StatusCode)
	require.Equal(t, ReasonMissingOwnership, ytErr.Reason)
	require.Equal(t, "The claimed asset is not owned by the caller.", ytErr.Description)
	require.Len(t, ytErr.Errors, 1)
	require.Equal(t, "youtubePartner.claims.update", ytErr.Errors[0].Domain)
	require.Equal(t, "missingOwnership: The claimed asset is not owned by the caller.", ytErr.Error())
}

// TestDecodeResponse_OAuthError verifies the OAuth-style error shape
// (string-valued "error") still decodes into ErrorType, with no Reason.
func TestDecodeResponse_OAuthError(t *testing.T) {
	body := `{"error":"invalid_grant","error_description":"Bad Request"}`

	err := DecodeResponse(errResponse(http.StatusBadRequest, body), nil)
	require.Error(t, err)

	var ytErr Error
	require.ErrorAs(t, err, &ytErr)
	require.Equal(t, http.StatusBadRequest, ytErr.StatusCode)
	require.Empty(t, ytErr.Reason)
	require.Equal(t, ErrTypeInvalidGrant, ytErr.ErrorType)
}

// TestDecodeResponse_Unparseable falls back to ErrTypeUnknown with the raw
// body preserved.
func TestDecodeResponse_Unparseable(t *testing.T) {
	body := `not json at all`

	err := DecodeResponse(errResponse(http.StatusInternalServerError, body), nil)
	require.Error(t, err)

	var ytErr Error
	require.ErrorAs(t, err, &ytErr)
	require.Equal(t, http.StatusInternalServerError, ytErr.StatusCode)
	require.Equal(t, ErrTypeUnknown, ytErr.ErrorType)
	require.Empty(t, ytErr.Reason)
	require.Equal(t, body, ytErr.Description)
}
