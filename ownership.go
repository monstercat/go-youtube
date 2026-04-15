package youtube

import (
	"net/http"
	"net/url"
)

const (
	OwnershipHistoryUrl = YoutubePartnerV1 + "/ownershipHistory"
)

func ownershipUrl(assetId string) string {
	return AssetsUrl + "/" + assetId + "/ownership"
}

// RightsOwnership represents the ownership data for an asset. It identifies an
// asset's owners and provides additional details about their ownership, such as
// the territories where they own the asset.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/ownership#RightsOwnership
type RightsOwnership struct {
	// General is a list that identifies the owners of an asset and the territories
	// where each owner has ownership. General asset ownership is used for all types
	// of assets and is the only type of ownership data that can be provided for
	// assets that are not compositions. Note: You cannot specify general ownership
	// rights and also specify either mechanical, performance, or synchronization
	// rights.
	General []*TerritoryOwners `json:"general,omitempty"`
	// Kind is the type of the API resource. For rightsOwnership resources, the
	// value is youtubePartner#rightsOwnership.
	Kind string `json:"kind,omitempty"`
	// Mechanical is a list that identifies owners of the mechanical rights for a
	// composition asset.
	Mechanical []*TerritoryOwners `json:"mechanical,omitempty"`
	// Performance is a list that identifies owners of the performance rights for a
	// composition asset.
	Performance []*TerritoryOwners `json:"performance,omitempty"`
	// Synchronization is a list that identifies owners of the synchronization
	// rights for a composition asset.
	Synchronization []*TerritoryOwners `json:"synchronization,omitempty"`
}

// RightsOwnershipHistory represents a set of ownership data provided for an asset.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/ownershipHistory#RightsOwnershipHistory
type RightsOwnershipHistory struct {
	// Kind is the type of the API resource. For ownership history resources, the
	// value is youtubePartner#rightsOwnershipHistory.
	Kind string `json:"kind,omitempty"`
	// Origination contains information that describes the metadata source.
	Origination *Origination `json:"origination,omitempty"`
	// Ownership contains the ownership data provided by the specified source
	// (origination) at the specified time (timeProvided).
	Ownership *RightsOwnership `json:"ownership,omitempty"`
	// TimeProvided is the time that the ownership data was provided.
	TimeProvided string `json:"timeProvided,omitempty"`
}

// OwnershipHistoryListResponse is the response for ownershipHistory.list.
type OwnershipHistoryListResponse struct {
	Items []*RightsOwnershipHistory `json:"items"`
	Kind  string                    `json:"kind,omitempty"`
}

// ── ownership.get ───────────────────────────────────────────────────────────

// GetOwnershipParams are parameters for ownership.get.
type GetOwnershipParams struct {
	// AssetId specifies the YouTube asset ID for which you are retrieving
	// ownership data.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetOwnershipParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetOwnership retrieves the ownership data provided for the specified asset by
// the content owner associated with the authenticated user.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/ownership/get
func GetOwnership(runner RequestRunner, p *GetOwnershipParams) (*RightsOwnership, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ownershipUrl(p.AssetId),
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out RightsOwnership
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── ownership.patch ─────────────────────────────────────────────────────────

// PatchOwnershipParams are parameters for ownership.patch.
type PatchOwnershipParams struct {
	// AssetId specifies the YouTube asset ID of the asset being updated.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	Ownership              *RightsOwnership
}

func (p *PatchOwnershipParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// PatchOwnership provides new ownership information for the specified asset.
// Note that YouTube may receive ownership information from multiple sources. For
// example, if an asset has multiple owners, each owner might send ownership data
// for the asset. YouTube algorithmically combines the ownership data received
// from all of those sources to generate the asset's canonical ownership data,
// which should provide the most comprehensive and accurate representation of the
// asset's ownership.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/ownership/patch
func PatchOwnership(runner RequestRunner, p *PatchOwnershipParams) (*RightsOwnership, error) {
	body, err := jsonBody(p.Ownership)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPatch,
		Url:    ownershipUrl(p.AssetId),
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out RightsOwnership
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── ownership.update ────────────────────────────────────────────────────────

// UpdateOwnershipParams are parameters for ownership.update.
type UpdateOwnershipParams struct {
	// AssetId specifies the YouTube asset ID of the asset being updated.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	Ownership              *RightsOwnership
}

func (p *UpdateOwnershipParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// UpdateOwnership provides new ownership information for the specified asset.
// Note that YouTube may receive ownership information from multiple sources. For
// example, if an asset has multiple owners, each owner might send ownership data
// for the asset. YouTube algorithmically combines the ownership data received
// from all of those sources to generate the asset's canonical ownership data,
// which should provide the most comprehensive and accurate representation of the
// asset's ownership.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/ownership/update
func UpdateOwnership(runner RequestRunner, p *UpdateOwnershipParams) (*RightsOwnership, error) {
	body, err := jsonBody(p.Ownership)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPut,
		Url:    ownershipUrl(p.AssetId),
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out RightsOwnership
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── ownershipHistory.list ───────────────────────────────────────────────────

// ListOwnershipHistoryParams are parameters for ownershipHistory.list.
type ListOwnershipHistoryParams struct {
	// AssetId specifies the YouTube asset ID of the asset for which you are
	// retrieving an ownership data history.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *ListOwnershipHistoryParams) Values() url.Values {
	v := url.Values{}
	if p.AssetId != "" {
		v.Set("assetId", p.AssetId)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// ListOwnershipHistory retrieves a list of the ownership data for an asset,
// regardless of which content owner provided the data. The list only includes
// the most recent ownership data for each content owner. However, if the content
// owner has submitted ownership data through multiple data sources (API, content
// feeds, etc.), the list will contain the most recent data for each content owner
// and data source.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/ownershipHistory/list
func ListOwnershipHistory(runner RequestRunner, p *ListOwnershipHistoryParams) (*OwnershipHistoryListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    OwnershipHistoryUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out OwnershipHistoryListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
