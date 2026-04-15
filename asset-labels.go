package youtube

import (
	"net/http"
	"net/url"
)

const (
	AssetLabelsUrl = YoutubePartnerV1 + "/assetLabels"
)

// AssetLabel represents a label that you can apply to an asset. Asset labels
// let you organize your assets into groups that you can search and filter.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetLabels#AssetLabel
type AssetLabel struct {
	// Kind identifies the API resource type. The value is "youtubePartner#assetLabel".
	Kind string `json:"kind,omitempty"`
	// LabelName is the label name. The value can contain up to 64 characters.
	LabelName string `json:"labelName,omitempty"`
}

// AssetLabelListResponse is the response for assetLabels.list.
type AssetLabelListResponse struct {
	// Items is a list of asset labels that match the request criteria.
	Items []*AssetLabel `json:"items"`
	// Kind identifies the API resource type. The value is "youtubePartner#assetLabelList".
	Kind string `json:"kind,omitempty"`
	// NextPageToken is the token that can be used as the value of the pageToken
	// parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`
	// PageInfo contains paging details for the result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
}

// ── assetLabels.insert ──────────────────────────────────────────────────────

// InsertAssetLabelParams are parameters for assetLabels.insert.
type InsertAssetLabelParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Label is the AssetLabel resource to insert.
	Label *AssetLabel
}

func (p *InsertAssetLabelParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// InsertAssetLabel inserts an asset label for an owner. The label can then
// be applied to assets to organize them into groups.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetLabels/insert
func InsertAssetLabel(runner RequestRunner, p *InsertAssetLabelParams) (*AssetLabel, error) {
	body, err := jsonBody(p.Label)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    AssetLabelsUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out AssetLabel
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assetLabels.list ────────────────────────────────────────────────────────

// ListAssetLabelsParams are parameters for assetLabels.list.
type ListAssetLabelsParams struct {
	// LabelPrefix filters for asset labels that have the specified prefix. For
	// example, a prefix of "GE" returns labels "GEO" and "GENRE".
	LabelPrefix string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// PageToken is the token that identifies a specific page in the result set
	// that should be returned. Set this to the value of nextPageToken from a
	// previous response to retrieve the next page.
	PageToken string
	// Q is a search query string. YouTube searches for labels that match the
	// query string.
	Q string
}

func (p *ListAssetLabelsParams) Values() url.Values {
	v := url.Values{}
	if p.LabelPrefix != "" {
		v.Set("labelPrefix", p.LabelPrefix)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.PageToken != "" {
		v.Set("pageToken", p.PageToken)
	}
	if p.Q != "" {
		v.Set("q", p.Q)
	}
	return v
}

// ListAssetLabels retrieves a list of all asset labels for an owner. Labels
// are used to organize assets into groups that can be searched and filtered.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetLabels/list
func ListAssetLabels(runner RequestRunner, p *ListAssetLabelsParams) (*AssetLabelListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    AssetLabelsUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out AssetLabelListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
