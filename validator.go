package youtube

import (
	"net/http"
	"net/url"
)

const (
	ValidatorUrl            = YoutubePartnerV1 + "/validator"
	ValidatorAsyncUrl       = YoutubePartnerV1 + "/validatorAsync"
	ValidatorAsyncStatusUrl = YoutubePartnerV1 + "/validatorAsyncStatus"
)

// ValidateRequest is the request body for validator.validate. It contains the
// metadata content to validate synchronously.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/validator/validate
type ValidateRequest struct {
	// Content is the metadata file content to validate. Typically XML in a
	// supported format (e.g., DDEX).
	Content string `json:"content,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#validateRequest".
	Kind string `json:"kind,omitempty"`
	// Locale is the locale of the metadata content (e.g., "en_US"). Used for
	// locale-specific validation rules and error messages.
	Locale string `json:"locale,omitempty"`
	// UploaderName is the name of the uploader associated with the content.
	UploaderName string `json:"uploaderName,omitempty"`
}

// ValidateResponse is the response for validator.validate. It contains the
// validation results including any errors found in the metadata content.
type ValidateResponse struct {
	// Errors is a list of validation errors found in the metadata content.
	Errors []*ValidateError `json:"errors,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#validateResponse".
	Kind string `json:"kind,omitempty"`
	// Status is the overall validation status (e.g., "ok" or "error").
	Status string `json:"status,omitempty"`
}

// ValidateAsyncRequest is the request body for validator.validateAsync. It
// contains the metadata content to validate asynchronously.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/validator/validateAsync
type ValidateAsyncRequest struct {
	// Content is the metadata file content to validate asynchronously.
	Content string `json:"content,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#validateAsyncRequest".
	Kind string `json:"kind,omitempty"`
	// UploaderName is the name of the uploader associated with the content.
	UploaderName string `json:"uploaderName,omitempty"`
}

// ValidateAsyncResponse is the response for validator.validateAsync. It
// contains a validation ID that can be used to poll for results.
type ValidateAsyncResponse struct {
	// Kind identifies the API resource type. The value is "youtubePartner#validateAsyncResponse".
	Kind string `json:"kind,omitempty"`
	// Status is the initial validation status (typically "processing").
	Status string `json:"status,omitempty"`
	// ValidationId is the unique ID assigned to this asynchronous validation
	// request. Use this ID with ValidateAsyncStatus to poll for results.
	ValidationId string `json:"validationId,omitempty"`
}

// ValidateStatusRequest is the request body for validator.validateAsyncStatus.
// It identifies a previously submitted asynchronous validation request.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/validator/validateAsyncStatus
type ValidateStatusRequest struct {
	// Kind identifies the API resource type. The value is "youtubePartner#validateStatusRequest".
	Kind string `json:"kind,omitempty"`
	// Locale is the locale for validation error messages (e.g., "en_US").
	Locale string `json:"locale,omitempty"`
	// ValidationId is the unique ID of the asynchronous validation request
	// returned by ValidateAsync.
	ValidationId string `json:"validationId,omitempty"`
}

// ValidateStatusResponse is the response for validator.validateAsyncStatus.
// It contains the current status and results of an asynchronous validation.
type ValidateStatusResponse struct {
	// Errors is a list of validation errors found in the metadata content.
	// Only populated when the validation is complete.
	Errors []*ValidateError `json:"errors,omitempty"`
	// IsMetadataOnly indicates whether the package contains only metadata
	// (no references or other binary content).
	IsMetadataOnly bool `json:"isMetadataOnly,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#validateStatusResponse".
	Kind string `json:"kind,omitempty"`
	// Status is the current validation status (e.g., "processing", "ok", "error").
	Status string `json:"status,omitempty"`
}

// ── validator.validate ──────────────────────────────────────────────────────

// ValidateParams are parameters for validator.validate.
type ValidateParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Request is the ValidateRequest containing the metadata content to validate.
	Request *ValidateRequest
}

func (p *ValidateParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// Validate validates a metadata file against YouTube's schema requirements.
// The validation is performed synchronously and the response contains any
// errors found. For large files, consider using ValidateAsync instead.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/validator/validate
func Validate(runner RequestRunner, p *ValidateParams) (*ValidateResponse, error) {
	body, err := jsonBody(p.Request)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    ValidatorUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out ValidateResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── validator.validateAsync ─────────────────────────────────────────────────

// ValidateAsyncParams are parameters for validator.validateAsync.
type ValidateAsyncParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Request is the ValidateAsyncRequest containing the metadata content to
	// validate asynchronously.
	Request *ValidateAsyncRequest
}

func (p *ValidateAsyncParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// ValidateAsync validates a metadata file asynchronously. The response returns
// a validationId that can be used with ValidateAsyncStatus to poll for the
// validation result. This is recommended for large metadata files.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/validator/validateAsync
func ValidateAsync(runner RequestRunner, p *ValidateAsyncParams) (*ValidateAsyncResponse, error) {
	body, err := jsonBody(p.Request)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    ValidatorAsyncUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out ValidateAsyncResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── validator.validateAsyncStatus ───────────────────────────────────────────

// ValidateAsyncStatusParams are parameters for validator.validateAsyncStatus.
type ValidateAsyncStatusParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Request is the ValidateStatusRequest identifying the async validation to check.
	Request *ValidateStatusRequest
}

func (p *ValidateAsyncStatusParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// ValidateAsyncStatus retrieves the current status and results of a previously
// submitted asynchronous validation request. Poll this method until the status
// is no longer "processing" to get the final validation results.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/validator/validateAsyncStatus
func ValidateAsyncStatus(runner RequestRunner, p *ValidateAsyncStatusParams) (*ValidateStatusResponse, error) {
	body, err := jsonBody(p.Request)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    ValidatorAsyncStatusUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out ValidateStatusResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
