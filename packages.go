package youtube

import (
	"net/http"
	"net/url"
)

const (
	PackageUrl = YoutubePartnerV1 + "/package"
)

// Package represents a content delivery package that can contain metadata,
// ownership, and match policy information for one or more assets.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/package#Package
type Package struct {
	// Content is the package's content, typically XML metadata in DDEX or
	// another supported format.
	Content string `json:"content,omitempty"`
	// CustomIds is a list of custom IDs associated with the assets in the package.
	CustomIds []string `json:"customIds,omitempty"`
	// Id is the unique ID that YouTube assigns to the package.
	Id string `json:"id,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#package".
	Kind string `json:"kind,omitempty"`
	// Locale is the locale of the package content (e.g., "en_US").
	Locale string `json:"locale,omitempty"`
	// Name is the package name.
	Name string `json:"name,omitempty"`
	// Status is the package's processing status (e.g., "processing", "succeeded", "failed").
	Status string `json:"status,omitempty"`
	// StatusReports is a list of status reports that describe the results of
	// processing the package. Each report may contain errors or warnings.
	StatusReports []*StatusReport `json:"statusReports,omitempty"`
	// TimeCreated is the date and time the package was created.
	TimeCreated string `json:"timeCreated,omitempty"`
	// Type is the package type (e.g., "package" or "localization").
	Type string `json:"type,omitempty"`
	// UploaderName is the name of the uploader who uploaded the package.
	UploaderName string `json:"uploaderName,omitempty"`
}

// PackageInsertResponse is the response for package.insert. It contains the
// processing status and any validation errors found during package insertion.
type PackageInsertResponse struct {
	// Errors is a list of errors found during validation of the package.
	Errors []*ValidateError `json:"errors,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#packageInsert".
	Kind string `json:"kind,omitempty"`
	// Resource is the Package resource that was inserted.
	Resource *Package `json:"resource,omitempty"`
	// Status is the insertion status (e.g., "succeeded", "failed", "processing").
	Status string `json:"status,omitempty"`
}

// ── package.get ─────────────────────────────────────────────────────────────

// GetPackageParams are parameters for package.get.
type GetPackageParams struct {
	// PackageId is the unique ID of the content delivery package to retrieve.
	PackageId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetPackageParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetPackage retrieves information for the specified content delivery package,
// including its processing status and any status reports.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/package/get
func GetPackage(runner RequestRunner, p *GetPackageParams) (*Package, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    PackageUrl + "/" + p.PackageId,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out Package
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── package.insert ──────────────────────────────────────────────────────────

// InsertPackageParams are parameters for package.insert.
type InsertPackageParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Package is the Package resource to insert. The content field should
	// contain the XML metadata body.
	Package *Package
}

func (p *InsertPackageParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// InsertPackage inserts a metadata-only package. The package can contain
// metadata, ownership, and match policy information for one or more assets.
// The response includes validation results for the package content.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/package/insert
func InsertPackage(runner RequestRunner, p *InsertPackageParams) (*PackageInsertResponse, error) {
	body, err := jsonBody(p.Package)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    PackageUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out PackageInsertResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
