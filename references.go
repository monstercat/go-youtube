package youtube

import (
	"net/http"
	"net/url"
	"strconv"
)

const (
	ReferencesUrl = YoutubePartnerV1 + "/references"
)

// Reference represents the actual content of an asset. YouTube compares newly
// uploaded videos to a library of references for the purpose of automatically
// generating claims for the asset's owner(s).
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/references#Reference
type Reference struct {
	// AssetId is the ID that uniquely identifies the asset that the reference is
	// associated with.
	AssetId string `json:"assetId,omitempty"`
	// AudioswapEnabled indicates that the reference content should be included in
	// YouTube's AudioSwap program when set to true.
	AudioswapEnabled bool `json:"audioswapEnabled,omitempty"`
	// ClaimId is present if the reference was created by associating an asset with
	// an existing YouTube video uploaded to a YouTube channel linked to your CMS
	// account. In that case, this field contains the ID of the claim representing
	// the resulting association between the asset and the video.
	ClaimId string `json:"claimId,omitempty"`
	// ContentType is the type of content that the reference represents.
	ContentType string `json:"contentType,omitempty"`
	// DuplicateLeader is the ID that uniquely identifies the reference that this
	// reference duplicates. This field is only present if the reference's status
	// is inactive with reason REASON_DUPLICATE_FOR_OWNERS.
	DuplicateLeader string `json:"duplicateLeader,omitempty"`
	// ExcludedIntervals is the list of time intervals from this reference that
	// will be ignored during the match process.
	ExcludedIntervals []*ExcludedInterval `json:"excludedIntervals,omitempty"`
	// FpDirect indicates that the reference is a pre-generated fingerprint when
	// set to true.
	FpDirect bool `json:"fpDirect,omitempty"`
	// HashCode is the MD5 hashcode of the reference content. Deprecated: this is
	// no longer populated.
	HashCode string `json:"hashCode,omitempty"`
	// Id is a value that YouTube assigns and uses to uniquely identify a
	// reference.
	Id string `json:"id,omitempty"`
	// IgnoreFpMatch indicates that the reference should not be used to generate
	// claims when set to true. This field is only used on AudioSwap references.
	IgnoreFpMatch bool `json:"ignoreFpMatch,omitempty"`
	// Kind is the type of the API resource. For reference resources, the value is
	// youtubePartner#reference.
	Kind string `json:"kind,omitempty"`
	// Length is the length of the reference in seconds.
	Length float64 `json:"length,omitempty"`
	// Origination contains information that describes the reference source.
	Origination *Origination `json:"origination,omitempty"`
	// Status is the reference's status.
	Status string `json:"status,omitempty"`
	// StatusReason is an explanation of how a reference entered its current state.
	// This value is only present if the reference's status is either inactive or
	// deleted.
	StatusReason string `json:"statusReason,omitempty"`
	// Urgent indicates that YouTube should prioritize Content ID processing for a
	// video file when set to true. YouTube processes urgent video files before
	// other files that are not marked as urgent. Note that marking all of your
	// files as urgent could delay processing for those files.
	Urgent bool `json:"urgent,omitempty"`
	// VideoId is present if the reference was created by associating an asset with
	// an existing YouTube video uploaded to a YouTube channel linked to your CMS
	// account. In that case, this field contains the ID of the source video.
	VideoId string `json:"videoId,omitempty"`
}

// ReferenceListResponse is the response for references.list.
type ReferenceListResponse struct {
	Items         []*Reference `json:"items"`
	Kind          string       `json:"kind,omitempty"`
	NextPageToken string       `json:"nextPageToken,omitempty"`
	PageInfo      *PageInfo    `json:"pageInfo,omitempty"`
}

// ── references.get ──────────────────────────────────────────────────────────

// GetReferenceParams are parameters for references.get.
type GetReferenceParams struct {
	// ReferenceId specifies the YouTube reference ID of the reference being
	// retrieved.
	ReferenceId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetReferenceParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetReference retrieves information about the specified reference.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/references/get
func GetReference(runner RequestRunner, p *GetReferenceParams) (*Reference, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ReferencesUrl + "/" + p.ReferenceId,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out Reference
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── references.insert ───────────────────────────────────────────────────────

// InsertReferenceParams are parameters for references.insert.
type InsertReferenceParams struct {
	// ClaimId specifies the YouTube claim ID of an existing claim from which a
	// reference should be created. (The claimed video is used as the reference
	// content.)
	ClaimId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	Reference              *Reference
}

func (p *InsertReferenceParams) Values() url.Values {
	v := url.Values{}
	if p.ClaimId != "" {
		v.Set("claimId", p.ClaimId)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// InsertReference creates a reference in one of the following ways: If your
// request is uploading a reference file, YouTube creates the reference from the
// provided content. You can provide either a video/audio file or a pre-generated
// fingerprint. If you are providing a pre-generated fingerprint, set the
// reference resource's fpDirect property to true in the request body. If you want
// to create a reference using a claimed video as the reference content, use the
// claimId parameter to identify the claim.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/references/insert
func InsertReference(runner RequestRunner, p *InsertReferenceParams) (*Reference, error) {
	body, err := jsonBody(p.Reference)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    ReferencesUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Reference
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── references.list ─────────────────────────────────────────────────────────

// ListReferencesParams are parameters for references.list.
type ListReferencesParams struct {
	// AssetId specifies the YouTube asset ID of the asset for which you are
	// retrieving references.
	AssetId string
	// Id specifies a comma-separated list of YouTube reference IDs to retrieve.
	Id string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	// PageToken specifies a token that identifies a particular page of results to
	// return. Set this parameter to the value of the nextPageToken value from the
	// previous API response to retrieve the next page of search results.
	PageToken string
}

func (p *ListReferencesParams) Values() url.Values {
	v := url.Values{}
	if p.AssetId != "" {
		v.Set("assetId", p.AssetId)
	}
	if p.Id != "" {
		v.Set("id", p.Id)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.PageToken != "" {
		v.Set("pageToken", p.PageToken)
	}
	return v
}

// ListReferences retrieves a list of references by ID or the list of references
// for the specified asset.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/references/list
func ListReferences(runner RequestRunner, p *ListReferencesParams) (*ReferenceListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ReferencesUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ReferenceListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── references.patch ────────────────────────────────────────────────────────

// PatchReferenceParams are parameters for references.patch.
type PatchReferenceParams struct {
	// ReferenceId specifies the YouTube reference ID of the reference being
	// updated.
	ReferenceId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	// ReleaseClaims indicates that you want to release all match claims associated
	// with this reference. This parameter only works when the claim's status is
	// being updated to 'inactive' - you can then set the parameter's value to true
	// to release all match claims produced by this reference.
	ReleaseClaims bool
	Reference     *Reference
}

func (p *PatchReferenceParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.ReleaseClaims {
		v.Set("releaseClaims", strconv.FormatBool(p.ReleaseClaims))
	}
	return v
}

// PatchReference patches a reference.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/references/patch
func PatchReference(runner RequestRunner, p *PatchReferenceParams) (*Reference, error) {
	body, err := jsonBody(p.Reference)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPatch,
		Url:    ReferencesUrl + "/" + p.ReferenceId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Reference
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── references.update ───────────────────────────────────────────────────────

// UpdateReferenceParams are parameters for references.update.
type UpdateReferenceParams struct {
	// ReferenceId specifies the YouTube reference ID of the reference being
	// updated.
	ReferenceId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	// ReleaseClaims indicates that you want to release all match claims associated
	// with this reference. This parameter only works when the claim's status is
	// being updated to 'inactive' - you can then set the parameter's value to true
	// to release all match claims produced by this reference.
	ReleaseClaims bool
	Reference     *Reference
}

func (p *UpdateReferenceParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.ReleaseClaims {
		v.Set("releaseClaims", strconv.FormatBool(p.ReleaseClaims))
	}
	return v
}

// UpdateReference updates a reference.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/references/update
func UpdateReference(runner RequestRunner, p *UpdateReferenceParams) (*Reference, error) {
	body, err := jsonBody(p.Reference)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPut,
		Url:    ReferencesUrl + "/" + p.ReferenceId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Reference
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
