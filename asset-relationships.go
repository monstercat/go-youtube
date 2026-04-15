package youtube

import (
	"net/http"
	"net/url"
)

const (
	AssetRelationshipsUrl = YoutubePartnerV1 + "/assetRelationships"
)

// AssetRelationship represents a relationship between two assets, where one
// is the parent and the other is the child. For example, a music video asset
// might be the child of a sound recording asset. Establishing these
// relationships enables YouTube to identify matching content more effectively.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetRelationships#AssetRelationship
type AssetRelationship struct {
	// ChildAssetId is the ID of the child (contained) asset.
	ChildAssetId string `json:"childAssetId,omitempty"`
	// Id is a server-assigned unique ID for the asset relationship.
	Id string `json:"id,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#assetRelationship".
	Kind string `json:"kind,omitempty"`
	// ParentAssetId is the ID of the parent (containing) asset.
	ParentAssetId string `json:"parentAssetId,omitempty"`
}

// AssetRelationshipListResponse is the response for assetRelationships.list.
type AssetRelationshipListResponse struct {
	// Items is a list of asset relationships that match the request criteria.
	Items []*AssetRelationship `json:"items"`
	// Kind identifies the API resource type. The value is "youtubePartner#assetRelationshipList".
	Kind string `json:"kind,omitempty"`
	// NextPageToken is the token that can be used as the value of the pageToken
	// parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`
	// PageInfo contains paging details for the result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
}

// ── assetRelationships.delete ───────────────────────────────────────────────

// DeleteAssetRelationshipParams are parameters for assetRelationships.delete.
type DeleteAssetRelationshipParams struct {
	// AssetRelationshipId is the ID of the asset relationship to delete.
	AssetRelationshipId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *DeleteAssetRelationshipParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// DeleteAssetRelationship deletes a relationship between two assets. Removing
// a relationship does not delete the assets themselves; it only severs the
// parent-child connection between them.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetRelationships/delete
func DeleteAssetRelationship(runner RequestRunner, p *DeleteAssetRelationshipParams) error {
	res, err := runner.Run(&Request{
		Method: http.MethodDelete,
		Url:    AssetRelationshipsUrl + "/" + p.AssetRelationshipId,
		Params: p.Values(),
	})
	if err != nil {
		return err
	}
	return DecodeResponse(res, nil)
}

// ── assetRelationships.insert ───────────────────────────────────────────────

// InsertAssetRelationshipParams are parameters for assetRelationships.insert.
type InsertAssetRelationshipParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Relationship is the AssetRelationship resource to insert.
	Relationship *AssetRelationship
}

func (p *InsertAssetRelationshipParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// InsertAssetRelationship creates a relationship between two assets. The
// relationship identifies one asset as the parent (containing) asset and the
// other as the child (contained) asset.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetRelationships/insert
func InsertAssetRelationship(runner RequestRunner, p *InsertAssetRelationshipParams) (*AssetRelationship, error) {
	body, err := jsonBody(p.Relationship)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    AssetRelationshipsUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out AssetRelationship
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assetRelationships.list ─────────────────────────────────────────────────

// ListAssetRelationshipsParams are parameters for assetRelationships.list.
type ListAssetRelationshipsParams struct {
	// AssetId is the asset ID of the asset for which you are retrieving
	// relationships. The API returns both parent and child relationships
	// associated with the specified asset.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// PageToken is the token that identifies a specific page in the result set
	// that should be returned. Set this to the value of nextPageToken from a
	// previous response to retrieve the next page.
	PageToken string
}

func (p *ListAssetRelationshipsParams) Values() url.Values {
	v := url.Values{}
	if p.AssetId != "" {
		v.Set("assetId", p.AssetId)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.PageToken != "" {
		v.Set("pageToken", p.PageToken)
	}
	return v
}

// ListAssetRelationships retrieves a list of relationships for a given asset.
// The list contains both parent and child relationships for the specified asset.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetRelationships/list
func ListAssetRelationships(runner RequestRunner, p *ListAssetRelationshipsParams) (*AssetRelationshipListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    AssetRelationshipsUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out AssetRelationshipListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
