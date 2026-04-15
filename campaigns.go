package youtube

import (
	"net/http"
	"net/url"
)

const (
	CampaignsUrl = YoutubePartnerV1 + "/campaigns"
)

// Campaign represents a YouTube advertising campaign used by content owners to
// promote their content. Each campaign associates ads with specific content.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/campaigns#Campaign
type Campaign struct {
	// CampaignData contains the campaign's data, including its promoted content,
	// timing, and source configuration.
	CampaignData *CampaignData `json:"campaignData,omitempty"`
	// Id is the unique ID that YouTube assigns to the campaign.
	Id string `json:"id,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#campaign".
	Kind string `json:"kind,omitempty"`
	// Status is the campaign's status. Valid values are "active", "paused", and "completed".
	Status string `json:"status,omitempty"`
	// TimeCreated is the date and time the campaign was created.
	TimeCreated string `json:"timeCreated,omitempty"`
	// TimeLastModified is the date and time the campaign was last modified.
	TimeLastModified string `json:"timeLastModified,omitempty"`
}

// CampaignData contains the core data for a campaign including its source,
// timing, and promoted content.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/campaigns#CampaignData
type CampaignData struct {
	// CampaignSource identifies the source of the campaign, such as a YouTube
	// channel or a content owner.
	CampaignSource *CampaignSource `json:"campaignSource,omitempty"`
	// ExpireTime is the date and time when the campaign expires. If not specified,
	// the campaign runs indefinitely.
	ExpireTime string `json:"expireTime,omitempty"`
	// Name is a user-defined name for the campaign.
	Name string `json:"name,omitempty"`
	// PromotedContent is the list of promoted content items associated with the campaign.
	PromotedContent []*PromotedContent `json:"promotedContent,omitempty"`
	// StartTime is the date and time when the campaign begins.
	StartTime string `json:"startTime,omitempty"`
}

// CampaignSource identifies the source of a campaign, such as a specific
// YouTube channel or content owner.
type CampaignSource struct {
	// SourceType is the type of source for the campaign (e.g., "contentOwner" or "channel").
	SourceType string `json:"sourceType,omitempty"`
	// SourceValue is a list of values that identify the campaign source. For
	// example, channel IDs or content owner IDs.
	SourceValue []string `json:"sourceValue,omitempty"`
}

// CampaignList is the response for campaigns.list.
type CampaignList struct {
	// Items is a list of campaigns that match the request criteria.
	Items []*Campaign `json:"items"`
	// Kind identifies the API resource type. The value is "youtubePartner#campaignList".
	Kind string `json:"kind,omitempty"`
}

// ── campaigns.get ───────────────────────────────────────────────────────────

// GetCampaignParams are parameters for campaigns.get.
type GetCampaignParams struct {
	// CampaignId is the unique ID of the campaign to retrieve.
	CampaignId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetCampaignParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetCampaign retrieves a specific campaign for an owner. The API response
// contains all data associated with the campaign, including its promoted
// content, timing, and source configuration.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/campaigns/get
func GetCampaign(runner RequestRunner, p *GetCampaignParams) (*Campaign, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    CampaignsUrl + "/" + p.CampaignId,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out Campaign
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── campaigns.list ──────────────────────────────────────────────────────────

// ListCampaignsParams are parameters for campaigns.list.
type ListCampaignsParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// PageToken is the token that identifies a specific page in the result set
	// that should be returned.
	PageToken string
}

func (p *ListCampaignsParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.PageToken != "" {
		v.Set("pageToken", p.PageToken)
	}
	return v
}

// ListCampaigns retrieves a list of campaigns for an owner. The list includes
// all campaigns associated with the authenticated content owner.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/campaigns/list
func ListCampaigns(runner RequestRunner, p *ListCampaignsParams) (*CampaignList, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    CampaignsUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out CampaignList
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── campaigns.insert ────────────────────────────────────────────────────────

// InsertCampaignParams are parameters for campaigns.insert.
type InsertCampaignParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Campaign is the Campaign resource to insert.
	Campaign *Campaign
}

func (p *InsertCampaignParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// InsertCampaign inserts a new campaign for an owner. The campaign associates
// ads with specific content using the provided campaign data.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/campaigns/insert
func InsertCampaign(runner RequestRunner, p *InsertCampaignParams) (*Campaign, error) {
	body, err := jsonBody(p.Campaign)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    CampaignsUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Campaign
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── campaigns.patch ─────────────────────────────────────────────────────────

// PatchCampaignParams are parameters for campaigns.patch.
type PatchCampaignParams struct {
	// CampaignId is the unique ID of the campaign to patch.
	CampaignId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Campaign is the Campaign resource with updated fields.
	Campaign *Campaign
}

func (p *PatchCampaignParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// PatchCampaign patches an existing campaign's data. This method supports
// patch semantics, meaning only the fields included in the request body will
// be updated; all other fields will retain their current values.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/campaigns/patch
func PatchCampaign(runner RequestRunner, p *PatchCampaignParams) (*Campaign, error) {
	body, err := jsonBody(p.Campaign)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPatch,
		Url:    CampaignsUrl + "/" + p.CampaignId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Campaign
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── campaigns.update ────────────────────────────────────────────────────────

// UpdateCampaignParams are parameters for campaigns.update.
type UpdateCampaignParams struct {
	// CampaignId is the unique ID of the campaign to update.
	CampaignId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Campaign is the complete Campaign resource to replace the existing campaign.
	Campaign *Campaign
}

func (p *UpdateCampaignParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// UpdateCampaign updates an existing campaign. This method replaces the entire
// campaign resource, so all fields must be provided. Use PatchCampaign for
// partial updates.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/campaigns/update
func UpdateCampaign(runner RequestRunner, p *UpdateCampaignParams) (*Campaign, error) {
	body, err := jsonBody(p.Campaign)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPut,
		Url:    CampaignsUrl + "/" + p.CampaignId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Campaign
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── campaigns.delete ────────────────────────────────────────────────────────

// DeleteCampaignParams are parameters for campaigns.delete.
type DeleteCampaignParams struct {
	// CampaignId is the unique ID of the campaign to delete.
	CampaignId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *DeleteCampaignParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// DeleteCampaign deletes a specified campaign for an owner. This permanently
// removes the campaign and its associated data.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/campaigns/delete
func DeleteCampaign(runner RequestRunner, p *DeleteCampaignParams) error {
	res, err := runner.Run(&Request{
		Method: http.MethodDelete,
		Url:    CampaignsUrl + "/" + p.CampaignId,
		Params: p.Values(),
	})
	if err != nil {
		return err
	}
	return DecodeResponse(res, nil)
}
