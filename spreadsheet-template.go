package youtube

import (
	"net/http"
	"net/url"
)

const (
	SpreadsheetTemplateUrl = YoutubePartnerV1 + "/spreadsheetTemplate"
)

// SpreadsheetTemplate represents a spreadsheet template that a content owner
// can use to bulk-manage assets, ownership, and match policies.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/spreadsheetTemplate#SpreadsheetTemplate
type SpreadsheetTemplate struct {
	// Kind identifies the API resource type. The value is "youtubePartner#spreadsheetTemplate".
	Kind string `json:"kind,omitempty"`
	// Status is the template's status.
	Status string `json:"status,omitempty"`
	// TemplateContent is the base64-encoded content of the spreadsheet template.
	TemplateContent string `json:"templateContent,omitempty"`
	// TemplateName is the display name of the template.
	TemplateName string `json:"templateName,omitempty"`
	// TemplateType is the type of spreadsheet template (e.g., "asset",
	// "ownership", "matchPolicy").
	TemplateType string `json:"templateType,omitempty"`
}

// SpreadsheetTemplateListResponse is the response for spreadsheetTemplate.list.
type SpreadsheetTemplateListResponse struct {
	// Items is a list of spreadsheet templates that match the request criteria.
	Items []*SpreadsheetTemplate `json:"items"`
	// Kind identifies the API resource type. The value is "youtubePartner#spreadsheetTemplateList".
	Kind string `json:"kind,omitempty"`
	// Status is the response status.
	Status string `json:"status,omitempty"`
}

// ── spreadsheetTemplate.list ────────────────────────────────────────────────

// ListSpreadsheetTemplatesParams are parameters for spreadsheetTemplate.list.
type ListSpreadsheetTemplatesParams struct {
	// Locale is the locale for the spreadsheet template content and header
	// labels (e.g., "en_US").
	Locale string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *ListSpreadsheetTemplatesParams) Values() url.Values {
	v := url.Values{}
	if p.Locale != "" {
		v.Set("locale", p.Locale)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// ListSpreadsheetTemplates retrieves a list of spreadsheet templates that a
// content owner can use to bulk-manage assets, ownership, and match policies.
// Templates are returned in the specified locale.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/spreadsheetTemplate/list
func ListSpreadsheetTemplates(runner RequestRunner, p *ListSpreadsheetTemplatesParams) (*SpreadsheetTemplateListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    SpreadsheetTemplateUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out SpreadsheetTemplateListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
