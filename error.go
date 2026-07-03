package youtube

import "errors"

var (
	ErrInvalidAccessType  = errors.New("invalid access type")
	ErrInvalidPart        = errors.New("invalid part")
	ErrInvalidPrompt      = errors.New("invalid prompt")
	ErrInvalidScope       = errors.New("invalid scope")
	ErrMissingClientId    = errors.New("missing client id")
	ErrMissingParts       = errors.New("missing parts")
	ErrMissingRedirectUri = errors.New("missing redirect uri")
	ErrMissingScopes      = errors.New("missing scopes")
	ErrNotFound           = errors.New("not found")
)

const (
	ErrAccessDenied        OAuthError = "access_denied"
	ErrAdminPolicyEnforced OAuthError = "admin_policy_enforced"
	ErrDisallowedUserAgent OAuthError = "disallowed_useragent"
	ErrOrgInternal         OAuthError = "org_internal"
	ErrRedirectUriMismatch OAuthError = "redirect_uri_mismatch"

	ErrTypeBody         ErrorType = "could not read body"
	ErrTypeInvalidGrant ErrorType = "invalid_grant"
	ErrTypeJSON         ErrorType = "json_error"
	ErrTypeUnknown      ErrorType = "unknown"
)

// Claim error reasons surfaced in a Google API error envelope's
// errors[].reason field (domain youtubePartner.claims.*). They are exposed
// via Error.Reason so callers can classify a claim failure by its cause.
// This is the claims-management subset of the Content ID API error
// reference; other domains (assets, policies, references, …) can be added
// as needed.
const (
	// claims.update / claims.patch
	ReasonAlreadyClaimed               = "alreadyClaimed"
	ReasonChannelMonetizationSuspended = "channelMonetizationSuspended"
	ReasonChannelNotActive             = "channelNotActive"
	ReasonClaimIsClosed                = "claimIsClosed"
	ReasonInvalidStatusTransition      = "invalidStatusTransition"
	ReasonMissingOwnership             = "missingOwnership"
	ReasonPolicyCannotBeChanged        = "policyCannotBeChanged"

	// claims.insert (in addition to the shared reasons above)
	ReasonExistingSoundRecordingOrMusicVideoClaim = "existingSoundRecordingOrMusicVideoClaim"
	ReasonInvalidContentOwner                     = "invalidContentOwner"
	ReasonPolicyTypeNotPermitted                  = "policyTypeNotPermitted"
	ReasonTakedownNotPermitted                    = "takedownNotPermitted"
	ReasonThirdPartyClaimNotAllowed               = "thirdPartyClaimNotAllowed"
	ReasonVideoIsPrivate                          = "videoIsPrivate"
	ReasonVideoNotOwned                           = "videoNotOwned"
	ReasonVideoNotProcessed                       = "videoNotProcessed"
	ReasonWrongContentType                        = "wrongContentType"

	// Common errors (any method)
	ReasonContentOwnerNotProvided  = "contentOwnerNotProvided"
	ReasonServiceAccountNotAllowed = "serviceAccountNotAllowed"
	ReasonInternalError            = "internalError"
	// ReasonInvalid is the global-domain rejection of an invalid request
	// parameter value; observed on claimSearch.list for an expired or
	// malformed pageToken.
	ReasonInvalid = "invalid"
)

type OAuthError string
type ErrorType string

func (e OAuthError) Error() string {
	return string(e)
}

// ErrorDetail is one entry from a Google API error envelope's errors[]
// array (e.g. {reason: "missingOwnership", domain: "youtubePartner.claims.update"}).
type ErrorDetail struct {
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// Error is the structured error returned for a failed YouTube API call.
// DecodeResponse assembles it from the response body — it is a domain
// value, not a JSON decode target, so it carries no struct tags (the two
// wire shapes are decoded into local types inside DecodeResponse).
type Error struct {
	// StatusCode is the HTTP status of the failed response.
	StatusCode int
	// ErrorType is the OAuth error code, or ErrTypeUnknown for a Content ID
	// API error (whose cause is carried in Reason instead).
	ErrorType ErrorType
	// Description is the human-readable message — the OAuth
	// error_description, or the Content ID envelope's error.message.
	Description string
	// Body is the raw response body.
	Body string
	// Reason is the first errors[].reason from a Content ID API error
	// envelope (e.g. "missingOwnership"); empty for OAuth-style errors.
	// Compare against the Reason* constants.
	Reason string
	// Errors holds every detail entry from a Content ID API error envelope.
	Errors []ErrorDetail
}

func (e Error) Error() string {
	if e.Reason != "" {
		return e.Reason + ": " + e.Description
	}
	return string(e.ErrorType) + ": " + e.Description
}
