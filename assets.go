package youtube

import (
	"net/http"
	"net/url"
)

const (
	AssetsUrl      = YoutubePartnerV1 + "/assets"
	AssetSearchUrl = YoutubePartnerV1 + "/assetSearch"
	AssetSharesUrl = YoutubePartnerV1 + "/assetShares"
)

// Asset represents a piece of intellectual property, such as a sound recording
// or television episode.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assets#Asset
type Asset struct {
	// AliasId is a list of asset IDs that can be used to refer to the asset.
	// The list contains values if the asset represents multiple constituent
	// assets that have been merged. In that case, any of the asset IDs
	// originally assigned to the constituent assets could be used to update
	// the master, or synthesized, asset.
	AliasId []string `json:"aliasId,omitempty"`
	// Id is an ID that YouTube assigns and uses to uniquely identify the asset.
	Id string `json:"id,omitempty"`
	// Kind is the type of the API resource. For asset resources, the value is
	// youtubePartner#asset.
	Kind string `json:"kind,omitempty"`
	// Label is a list of asset labels on the asset.
	Label []string `json:"label,omitempty"`
	// Licensability contains asset licensability information.
	Licensability *AssetLicensability `json:"licensability,omitempty"`
	// MatchPolicy contains information about the asset's match policy, which
	// YouTube applies to user-uploaded videos that match the asset.
	MatchPolicy *AssetMatchPolicy `json:"matchPolicy,omitempty"`
	// MatchPolicyEffective is the effective match policy for the asset.
	MatchPolicyEffective *AssetMatchPolicy `json:"matchPolicyEffective,omitempty"`
	// MatchPolicyMine is the match policy set by the authenticated user's
	// content owner.
	MatchPolicyMine *AssetMatchPolicy `json:"matchPolicyMine,omitempty"`
	// Metadata contains information that identifies and describes the asset.
	// This information could be used to search for the asset or to eliminate
	// duplication within YouTube's database.
	Metadata *Metadata `json:"metadata,omitempty"`
	// MetadataEffective is the effective metadata for the asset.
	MetadataEffective *Metadata `json:"metadataEffective,omitempty"`
	// MetadataMine is the metadata provided by the authenticated user's
	// content owner.
	MetadataMine *Metadata `json:"metadataMine,omitempty"`
	// NWayRevenueSharing contains N way revenue sharing (Pangea) information.
	NWayRevenueSharing *NWayRevenueSharing `json:"nWayRevenueSharing,omitempty"`
	// Ownership identifies an asset's owners and provides additional details
	// about their ownership, such as the territories where they own the asset.
	Ownership *RightsOwnership `json:"ownership,omitempty"`
	// OwnershipConflicts contains information about the asset's ownership
	// conflicts.
	OwnershipConflicts *OwnershipConflicts `json:"ownershipConflicts,omitempty"`
	// OwnershipEffective is the effective ownership for the asset.
	OwnershipEffective *RightsOwnership `json:"ownershipEffective,omitempty"`
	// OwnershipMine is the ownership data provided by the authenticated user's
	// content owner.
	OwnershipMine *RightsOwnership `json:"ownershipMine,omitempty"`
	// Status is the asset's status.
	Status string `json:"status,omitempty"`
	// TimeCreated is the date and time the asset was created. The value is
	// specified in RFC 3339 (YYYY-MM-DDThh:mm:ss.000Z) format.
	TimeCreated string `json:"timeCreated,omitempty"`
	// Type is the asset's type. This value determines the metadata fields that
	// you can set for the asset. In addition, certain API functions may only be
	// supported for specific types of assets.
	Type string `json:"type,omitempty"`
}

// AssetListResponse is the response for assets.list.
type AssetListResponse struct {
	Items []*Asset `json:"items"`
	Kind  string   `json:"kind,omitempty"`
}

// AssetSnippet is a compact representation of an asset returned by assetSearch.
// Each search result contains a snippet that provides key details about an
// asset. Note that the asset metadata returned in search results may be
// slightly out of date.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetSearch#AssetSnippet
type AssetSnippet struct {
	// CustomId is the custom ID assigned by the content owner to this asset.
	CustomId string `json:"customId,omitempty"`
	// Id is an ID that YouTube assigns and uses to uniquely identify the asset.
	Id string `json:"id,omitempty"`
	// Isrc is the ISRC (International Standard Recording Code) for this asset.
	Isrc string `json:"isrc,omitempty"`
	// Isrcs is the list of ISRCs (International Standard Recording Code) for
	// this asset.
	Isrcs []string `json:"isrcs,omitempty"`
	// Iswc is the ISWC (International Standard Musical Work Code) for this
	// asset.
	Iswc string `json:"iswc,omitempty"`
	// Iswcs is the list of ISWCs (International Standard Musical Work Code)
	// for this asset.
	Iswcs []string `json:"iswcs,omitempty"`
	// Kind is the type of the API resource. For this operation, the value is
	// youtubePartner#assetSnippet.
	Kind string `json:"kind,omitempty"`
	// TimeCreated is the date and time the asset was created. The value is
	// specified in RFC 3339 (YYYY-MM-DDThh:mm:ss.000Z) format.
	TimeCreated string `json:"timeCreated,omitempty"`
	// Title is the title of this asset.
	Title string `json:"title,omitempty"`
	// Type is the asset's type. This value determines which metadata fields
	// might be included in the metadata object.
	Type string `json:"type,omitempty"`
}

// AssetSearchResponse is the response for assetSearch.list.
type AssetSearchResponse struct {
	Items         []*AssetSnippet `json:"items"`
	Kind          string          `json:"kind,omitempty"`
	NextPageToken string          `json:"nextPageToken,omitempty"`
	PageInfo      *PageInfo       `json:"pageInfo,omitempty"`
}

// AssetShare identifies a relationship between two representations of an asset
// resource (a composition view and a composition share), which is only relevant
// to composition assets.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetShares#AssetShare
type AssetShare struct {
	// Kind is the type of the API resource. For this resource, the value is
	// youtubePartner#assetShare.
	Kind string `json:"kind,omitempty"`
	// ShareId is a value that YouTube assigns and uses to uniquely identify the
	// asset share.
	ShareId string `json:"shareId,omitempty"`
	// ViewId is a value that YouTube assigns and uses to uniquely identify the
	// asset view.
	ViewId string `json:"viewId,omitempty"`
}

// AssetShareListResponse is the response for assetShares.list.
type AssetShareListResponse struct {
	Items         []*AssetShare `json:"items"`
	Kind          string        `json:"kind,omitempty"`
	NextPageToken string        `json:"nextPageToken,omitempty"`
	PageInfo      *PageInfo     `json:"pageInfo,omitempty"`
}

// ── assets.get ──────────────────────────────────────────────────────────────

// GetAssetParams are parameters for assets.get.
type GetAssetParams struct {
	// AssetId specifies the YouTube asset ID of the asset being retrieved.
	AssetId string
	// FetchMatchPolicy specifies a comma-separated list of versions of the
	// asset's match policy that should be returned in the API response.
	// Acceptable values are: "effective", "mine", "none" (default).
	FetchMatchPolicy string
	// FetchMetadata specifies the version of the asset's metadata that should
	// be returned in the API response. In some cases, YouTube receives metadata
	// for an asset from multiple sources, such as when different partners own
	// the asset in different territories.
	FetchMetadata string
	// FetchOwnership specifies a comma-separated list of versions of the
	// asset's ownership data that should be returned in the API response.
	// Acceptable values are: "effective", "mine", "none" (default).
	FetchOwnership string
	// FetchOwnershipConflicts allows you to retrieve information about
	// ownership conflicts.
	FetchOwnershipConflicts bool
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetAssetParams) Values() url.Values {
	v := url.Values{}
	if p.FetchMatchPolicy != "" {
		v.Set("fetchMatchPolicy", p.FetchMatchPolicy)
	}
	if p.FetchMetadata != "" {
		v.Set("fetchMetadata", p.FetchMetadata)
	}
	if p.FetchOwnership != "" {
		v.Set("fetchOwnership", p.FetchOwnership)
	}
	if p.FetchOwnershipConflicts {
		v.Set("fetchOwnershipConflicts", "true")
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetAsset retrieves the metadata for the specified asset. Note that if the
// request identifies an asset that has been merged with another asset, meaning
// that YouTube identified the requested asset as a duplicate, then the request
// retrieves the merged, or synthesized, asset.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assets/get
func GetAsset(runner RequestRunner, p *GetAssetParams) (*Asset, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    AssetsUrl + "/" + p.AssetId,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out Asset
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assets.list ─────────────────────────────────────────────────────────────

// ListAssetsParams are parameters for assets.list.
type ListAssetsParams struct {
	// Id specifies a comma-separated list of YouTube Asset IDs that identify
	// the assets you want to retrieve. If you try to retrieve an asset that
	// YouTube identified as a duplicate and merged with another asset, the API
	// response only returns the synthesized asset.
	Id string
	// FetchMatchPolicy specifies a comma-separated list of versions of the
	// asset's match policy that should be returned in the API response.
	// Acceptable values are: "effective", "mine", "none" (default).
	FetchMatchPolicy string
	// FetchMetadata specifies the version of the asset's metadata that should
	// be returned in the API response. In some cases, YouTube receives metadata
	// for an asset from multiple sources, such as when different partners own
	// the asset in different territories.
	FetchMetadata string
	// FetchOwnership specifies a comma-separated list of versions of the
	// asset's ownership data that should be returned in the API response.
	// Acceptable values are: "effective", "mine", "none" (default).
	FetchOwnership string
	// FetchOwnershipConflicts allows you to retrieve information about
	// ownership conflicts.
	FetchOwnershipConflicts bool
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *ListAssetsParams) Values() url.Values {
	v := url.Values{}
	if p.Id != "" {
		v.Set("id", p.Id)
	}
	if p.FetchMatchPolicy != "" {
		v.Set("fetchMatchPolicy", p.FetchMatchPolicy)
	}
	if p.FetchMetadata != "" {
		v.Set("fetchMetadata", p.FetchMetadata)
	}
	if p.FetchOwnership != "" {
		v.Set("fetchOwnership", p.FetchOwnership)
	}
	if p.FetchOwnershipConflicts {
		v.Set("fetchOwnershipConflicts", "true")
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// ListAssets retrieves a list of assets based on asset metadata. The method can
// retrieve all assets or only assets owned by the content owner. Note that in
// cases where duplicate assets have been merged, the API response only contains
// the synthesized asset. (It does not contain the constituent assets that were
// merged into the synthesized asset.)
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assets/list
func ListAssets(runner RequestRunner, p *ListAssetsParams) (*AssetListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    AssetsUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out AssetListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assets.insert ───────────────────────────────────────────────────────────

// InsertAssetParams are parameters for assets.insert.
type InsertAssetParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Asset is the asset resource to insert.
	Asset *Asset
}

func (p *InsertAssetParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// InsertAsset inserts an asset with the specified metadata. After inserting an
// asset, you can set its ownership data and match policy.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assets/insert
func InsertAsset(runner RequestRunner, p *InsertAssetParams) (*Asset, error) {
	body, err := jsonBody(p.Asset)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    AssetsUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Asset
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assets.patch ────────────────────────────────────────────────────────────

// PatchAssetParams are parameters for assets.patch.
type PatchAssetParams struct {
	// AssetId specifies the YouTube asset ID of the asset being patched.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Asset is the asset resource with updated fields.
	Asset *Asset
}

func (p *PatchAssetParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// PatchAsset patches the metadata for the specified asset.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assets/patch
func PatchAsset(runner RequestRunner, p *PatchAssetParams) (*Asset, error) {
	body, err := jsonBody(p.Asset)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPatch,
		Url:    AssetsUrl + "/" + p.AssetId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Asset
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assets.update ───────────────────────────────────────────────────────────

// UpdateAssetParams are parameters for assets.update.
type UpdateAssetParams struct {
	// AssetId specifies the YouTube asset ID of the asset being updated.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Asset is the asset resource with updated fields.
	Asset *Asset
}

func (p *UpdateAssetParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// UpdateAsset updates the metadata for the specified asset.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assets/update
func UpdateAsset(runner RequestRunner, p *UpdateAssetParams) (*Asset, error) {
	body, err := jsonBody(p.Asset)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPut,
		Url:    AssetsUrl + "/" + p.AssetId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Asset
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assetSearch.list ────────────────────────────────────────────────────────

// SearchAssetsParams are parameters for assetSearch.list.
type SearchAssetsParams struct {
	// CreatedAfter restricts the set of returned assets to ones originally
	// created on or after the specified datetime. For example:
	// 2015-01-29T23:00:00Z
	CreatedAfter string
	// CreatedBefore restricts the set of returned assets to ones originally
	// created on or before the specified datetime. For example:
	// 2015-01-29T23:00:00Z
	CreatedBefore string
	// HasConflicts enables you to only retrieve assets that have ownership
	// conflicts. The only valid value is true.
	HasConflicts bool
	// IncludeAnyProvidedLabel, if set to true, will search for assets that
	// contain any of the provided labels; otherwise will search for assets
	// that contain all the provided labels.
	IncludeAnyProvidedLabel bool
	// Isrcs is a comma-separated list of up to 50 ISRCs. If you specify a
	// value for this parameter, the API server ignores any values set for q,
	// includeAnyProvidedLabel, hasConflicts, labels, metadataSearchFields,
	// sort, and type.
	Isrcs string
	// Labels specifies the assets with certain asset labels that you want to
	// retrieve. The parameter value is a comma-separated list of asset labels.
	Labels string
	// MetadataSearchFields specifies which metadata fields to search by. It is
	// a comma-separated list of metadata field and value pairs connected by
	// colon. For example: customId:my_custom_id,artist:Dandexx
	MetadataSearchFields string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// OwnershipRestriction specifies the ownership filtering option for the
	// search. By default the search is performed in the assets owned by the
	// currently authenticated user only.
	OwnershipRestriction string
	// PageToken specifies a token that identifies a particular page of results
	// to return. Set this parameter to the value of the nextPageToken value
	// from the previous API response to retrieve the next page of search
	// results.
	PageToken string
	// Q is the query string. YouTube searches within the id, type, and
	// customId fields for all assets as well as in numerous other metadata
	// fields that vary for different types of assets.
	Q string
	// Sort specifies how the search results should be sorted. Note that
	// results are always sorted in descending order.
	Sort string
	// Type specifies the types of assets that you want to retrieve. The
	// parameter value is a comma-separated list of asset types.
	Type string
}

func (p *SearchAssetsParams) Values() url.Values {
	v := url.Values{}
	if p.CreatedAfter != "" {
		v.Set("createdAfter", p.CreatedAfter)
	}
	if p.CreatedBefore != "" {
		v.Set("createdBefore", p.CreatedBefore)
	}
	if p.HasConflicts {
		v.Set("hasConflicts", "true")
	}
	if p.IncludeAnyProvidedLabel {
		v.Set("includeAnyProvidedlabel", "true")
	}
	if p.Isrcs != "" {
		v.Set("isrcs", p.Isrcs)
	}
	if p.Labels != "" {
		v.Set("labels", p.Labels)
	}
	if p.MetadataSearchFields != "" {
		v.Set("metadataSearchFields", p.MetadataSearchFields)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.OwnershipRestriction != "" {
		v.Set("ownershipRestriction", p.OwnershipRestriction)
	}
	if p.PageToken != "" {
		v.Set("pageToken", p.PageToken)
	}
	if p.Q != "" {
		v.Set("q", p.Q)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	if p.Type != "" {
		v.Set("type", p.Type)
	}
	return v
}

// SearchAssets searches for assets based on asset metadata. The method can
// retrieve all assets or only assets owned by the content owner. This method
// mimics the functionality of the advanced search feature on the Assets page
// in CMS.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetSearch/list
func SearchAssets(runner RequestRunner, p *SearchAssetsParams) (*AssetSearchResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    AssetSearchUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out AssetSearchResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assetShares.list ────────────────────────────────────────────────────────

// ListAssetSharesParams are parameters for assetShares.list.
type ListAssetSharesParams struct {
	// AssetId specifies the asset ID for which you are retrieving data. The
	// parameter can be an asset view ID or an asset share ID. If the value is
	// an asset view ID, the API response identifies any asset share IDs mapped
	// to the asset view. If the value is an asset share ID, the API response
	// identifies any asset view IDs that map to that asset share.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// PageToken specifies a token that identifies a particular page of results
	// to return. Set this parameter to the value of the nextPageToken value
	// from the previous API response to retrieve the next page of search
	// results.
	PageToken string
}

func (p *ListAssetSharesParams) Values() url.Values {
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

// ListAssetShares either retrieves a list of asset shares the partner owns and
// that map to a specified asset view ID or it retrieves a list of asset views
// associated with a specified asset share ID owned by the partner.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetShares/list
func ListAssetShares(runner RequestRunner, p *ListAssetSharesParams) (*AssetShareListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    AssetSharesUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out AssetShareListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
