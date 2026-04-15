package youtube

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var (
	ErrInvalidClaimSearchParams = errors.New("invalid claim search params")
	ErrInvalidPatchClaimsParams = errors.New("invalid patch claims params")
	ErrMissingClaimId           = errors.New("claim id is required")
)

// ClaimStatus - status of a claim.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claims#Claim.FIELDS.status
type ClaimStatus string

const (
	ClaimsUrl       = YoutubePartnerV1 + "/claims"
	SearchClaimsUrl = YoutubePartnerV1 + "/claimSearch"
	ClaimHistoryUrl = YoutubePartnerV1 + "/claimHistory"

	ClaimStatusActive          ClaimStatus = "active"
	ClaimStatusAppealed        ClaimStatus = "appealed"
	ClaimStatusDisputed        ClaimStatus = "disputed"
	ClaimStatusInactive        ClaimStatus = "inactive"
	ClaimStatusPending         ClaimStatus = "pending"
	ClaimStatusPotential       ClaimStatus = "potential"
	ClaimStatusRoutedForReview ClaimStatus = "routedForReview"
	ClaimStatusTakedown        ClaimStatus = "takedown"

	ClaimSortDate      = "date"
	ClaimSortViewCount = "viewCount"
)

var (
	DateRegexp = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
)

func (s ClaimStatus) Valid() bool {
	switch s {
	case ClaimStatusActive,
		ClaimStatusAppealed,
		ClaimStatusDisputed,
		ClaimStatusInactive,
		ClaimStatusPending,
		ClaimStatusPotential,
		ClaimStatusRoutedForReview,
		ClaimStatusTakedown:
		return true
	}
	return false
}

// Claim represents a YouTube Content ID claim resource.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claims#Claim
type Claim struct {
	// AppliedPolicy: The applied policy for the claim.
	AppliedPolicy *Policy `json:"appliedPolicy,omitempty"`

	// AssetId: The unique YouTube asset ID that identifies the asset
	// associated with the claim.
	AssetId string `json:"assetId,omitempty"`

	// BlockOutsideOwnership: Indicates whether or not the claimed video
	// should be blocked anywhere it is not explicitly owned.
	BlockOutsideOwnership bool `json:"blockOutsideOwnership,omitempty"`

	// ContentType: This value indicates whether the claim covers the audio,
	// video, or audiovisual portion of the claimed content.
	ContentType string `json:"contentType,omitempty"`

	// Id: The ID that YouTube assigns and uses to uniquely identify the
	// claim.
	Id string `json:"id,omitempty"`

	// IsPartnerUploaded: Indicates whether or not the claim is a partner
	// uploaded claim.
	IsPartnerUploaded bool `json:"isPartnerUploaded,omitempty"`

	// Kind: The type of the API resource. For claim resources, this value
	// is youtubePartner#claim.
	Kind string `json:"kind,omitempty"`

	// MatchInfo: Match information about the claim.
	MatchInfo *MatchInfo `json:"matchInfo,omitempty"`

	// Origin: The origin of the claim.
	Origin *Origin `json:"origin,omitempty"`

	// Policy: The policy provided by the viewer or channel owner for the
	// claimed video.
	Policy *Policy `json:"policy,omitempty"`

	// Status: The claim's status.
	Status ClaimStatus `json:"status,omitempty"`

	// StudioInfo: Contains URLs linking to the YouTube Studio claim page.
	StudioInfo *StudioInfo `json:"studioInfo,omitempty"`

	// TimeCreated: The time the claim was created.
	TimeCreated string `json:"timeCreated,omitempty"`

	// TimeStatusLastModified: The time that the claim's status was last modified.
	TimeStatusLastModified string `json:"timeStatusLastModified"`

	// UgcType: The type of UGC content (e.g., "song").
	UgcType string `json:"ugcType,omitempty"`

	// VideoId: The unique YouTube video ID that identifies the video
	// associated with the claim.
	VideoId string `json:"videoId,omitempty"`
}

// ClaimSnippet is a compact representation of a claim returned by claimSearch.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claimSearch#ClaimSnippet
type ClaimSnippet struct {
	// AssetId: The unique YouTube asset ID that identifies the asset associated
	// with the claim.
	AssetId string `json:"assetId,omitempty"`

	// ContentType: This value indicates whether the claim covers the audio,
	// video, or audiovisual portion of the claimed content.
	ContentType string `json:"contentType,omitempty"`

	// EngagedViews: The number of engaged views on the claimed video.
	EngagedViews string `json:"engagedViews,omitempty"`

	// Id: The ID that YouTube assigns and uses to uniquely identify the claim.
	Id string `json:"id,omitempty"`

	// IsPartnerUploaded: Indicates whether or not the claim is a partner
	// uploaded claim.
	IsPartnerUploaded bool `json:"isPartnerUploaded,omitempty"`

	// IsVideoShortsEligible: Indicates whether the claimed video is eligible
	// for Shorts.
	IsVideoShortsEligible bool `json:"isVideoShortsEligible,omitempty"`

	// Kind: The type of the API resource. For claimSnippet resources, this
	// value is youtubePartner#claimSnippet.
	Kind string `json:"kind,omitempty"`

	// Origin: The origin of the claim.
	Origin *Origin `json:"origin,omitempty"`

	// Status: The claim's status.
	Status string `json:"status,omitempty"`

	// StudioInfo: Contains URLs linking to the YouTube Studio claim page.
	StudioInfo *StudioInfo `json:"studioInfo,omitempty"`

	// ThirdPartyClaim: Indicates whether the claim is a third-party claim.
	ThirdPartyClaim bool `json:"thirdPartyClaim,omitempty"`

	// TimeCreated: The time the claim was created.
	TimeCreated string `json:"timeCreated,omitempty"`

	// TimeStatusLastModified: The time that the claim's status was last
	// modified.
	TimeStatusLastModified string `json:"timeStatusLastModified,omitempty"`

	// VideoId: The unique YouTube video ID that identifies the video associated
	// with the claim.
	VideoId string `json:"videoId,omitempty"`

	// VideoTitle: The title of the claimed video.
	VideoTitle string `json:"videoTitle,omitempty"`

	// VideoViews: The number of views on the claimed video.
	VideoViews string `json:"videoViews,omitempty"`
}

// ClaimSearchResponse is the response for claimSearch.list. It contains a list
// of claims that match the search criteria.
type ClaimSearchResponse struct {
	// Items: A list of claims that match the request criteria.
	Items []*ClaimSnippet `json:"items"`

	// Kind: The type of the API resource. For this response, the value is
	// youtubePartner#claimSearchResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the pageToken
	// parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo"`

	// PreviousPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PreviousPageToken string `json:"previousPageToken,omitempty"`
}

// ClaimListResponse is the response for claims.list. It contains a list of
// claims administered by the content owner associated with the authenticated
// user.
type ClaimListResponse struct {
	// Items: A list of claims that match the request criteria.
	Items []*Claim `json:"items"`

	// Kind: The type of the API resource. For this response, the value is
	// youtubePartner#claimListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the pageToken
	// parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PageInfo: The pageInfo object encapsulates paging information for the
	// result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// PreviousPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the previous page in the result set.
	PreviousPageToken string `json:"previousPageToken,omitempty"`
}

// ClaimEvent represents an event in a claim's history.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claimHistory#ClaimEvent
type ClaimEvent struct {
	// Kind: The type of the API resource. For claimEvent resources, this value
	// is youtubePartner#claimEvent.
	Kind string `json:"kind,omitempty"`

	// Reason: The reason the claim status changed. This field is only populated
	// when the type is statusChange.
	Reason string `json:"reason,omitempty"`

	// Source: The source information for the event.
	Source *Source `json:"source,omitempty"`

	// Time: The time when the event occurred.
	Time string `json:"time,omitempty"`

	// Type: The type of the claim event.
	Type string `json:"type,omitempty"`

	// TypeDetails: Additional details about the event type.
	TypeDetails *TypeDetails `json:"typeDetails,omitempty"`
}

// ClaimHistory is the response for claimHistory.get.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claimHistory#ClaimHistory
type ClaimHistory struct {
	// Event: A list of claim history events.
	Event []*ClaimEvent `json:"event,omitempty"`

	// Id: The ID that YouTube assigns and uses to uniquely identify the claim
	// history.
	Id string `json:"id,omitempty"`

	// Kind: The type of the API resource. For claimHistory resources, this
	// value is youtubePartner#claimHistory.
	Kind string `json:"kind,omitempty"`

	// UploaderChannelId: The unique YouTube channel ID of the video uploader.
	UploaderChannelId string `json:"uploaderChannelId,omitempty"`
}

// ── claimSearch.list ────────────────────────────────────────────────────────

// SearchClaimsParams are parameters for the claimSearch.list method. Exactly one
// of AssetId, Q, ReferenceId, VideoIds, or Status must be set. The API returns
// a paginated list of ClaimSnippet resources.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claimSearch/list
type SearchClaimsParams struct {
	// AssetId: The assetId parameter specifies the YouTube asset ID of the
	// asset for which you are retrieving claims.
	AssetId string

	// ContentType: The contentType parameter specifies the content type of
	// claims that you want to retrieve.
	ContentType string

	// CreatedAfter: The createdAfter parameter allows you to restrict the set
	// of returned claims to ones created on or after the specified date
	// (inclusive).
	CreatedAfter string

	// CreatedBefore: The createdBefore parameter allows you to restrict the set
	// of returned claims to ones created before the specified date (exclusive).
	CreatedBefore string

	// InactiveReasons: The inactiveReasons parameter allows you to specify what
	// kind of inactive claims you want to find. This is a comma-separated list.
	InactiveReasons string

	// IncludeThirdPartyClaims: Used along with the videoId parameter, this
	// parameter determines whether or not to include third-party claims in the
	// search results.
	IncludeThirdPartyClaims bool

	// IsVideoShortsEligible: The isVideoShortsEligible parameter filters search
	// results to only include claims on videos that are eligible for YouTube
	// Shorts.
	IsVideoShortsEligible bool

	// OnBehalfOfContentOwner: The onBehalfOfContentOwner parameter identifies
	// the content owner that the user is acting on behalf of. This parameter
	// supports users whose accounts are associated with multiple content owners.
	OnBehalfOfContentOwner string

	// Origin: The origin parameter restricts the set of search results to only
	// include claims with the specified origin.
	Origin string

	// PageToken: The pageToken parameter specifies a token that identifies a
	// particular page of results to return. For example, set this parameter to
	// the value of the nextPageToken from the previous API response to retrieve
	// the next page of search results.
	PageToken string

	// PartnerUploaded: The partnerUploaded parameter specifies whether or not
	// to restrict the search results to only partner-uploaded claims.
	PartnerUploaded bool

	// Q: The q parameter specifies the query string to use to filter search
	// results. YouTube searches for the query string in the video title and
	// video description of claimed videos.
	Q string

	// ReferenceId: The referenceId parameter specifies the YouTube reference ID
	// of the reference for which you are retrieving claims.
	ReferenceId string

	// Sort: The sort parameter specifies the method that will be used to order
	// resources in the API response. Acceptable values are "date" and
	// "viewCount".
	Sort string

	// Status: The status parameter restricts your results to only claims in the
	// specified status.
	Status ClaimStatus

	// StatusModifiedAfter: The statusModifiedAfter parameter allows you to
	// restrict the result set to only include claims that have had their status
	// modified on or after the specified date (inclusive). The date specified
	// must be on or after June 30, 2016 (2016-06-30). The date format is
	// YYYY-MM-DD.
	StatusModifiedAfter string

	// StatusModifiedBefore: The statusModifiedBefore parameter allows you to
	// restrict the result set to only include claims that have had their status
	// modified before the specified date (exclusive). The date format is
	// YYYY-MM-DD.
	StatusModifiedBefore string

	// VideoIds: The videoId parameter specifies a comma-separated list of
	// YouTube video IDs for which you are retrieving claims.
	VideoIds []string
}

func (p *SearchClaimsParams) Validate() bool {
	required := []string{
		p.AssetId,
		p.Q,
		p.ReferenceId,
		strings.Join(p.VideoIds, ","),
		string(p.Status),
	}
	count := 0
	for _, r := range required {
		if r != "" {
			count++
		}
		if count > 1 {
			return false
		}
	}
	if p.Status != "" && !p.Status.Valid() {
		return false
	}
	return count > 0
}

func (p *SearchClaimsParams) Values() url.Values {
	vals := url.Values{}
	if p.AssetId != "" {
		vals.Set("assetId", p.AssetId)
	}
	if p.ContentType != "" {
		vals.Set("contentType", p.ContentType)
	}
	if p.CreatedAfter != "" {
		vals.Set("createdAfter", p.CreatedAfter)
	}
	if p.CreatedBefore != "" {
		vals.Set("createdBefore", p.CreatedBefore)
	}
	if p.InactiveReasons != "" {
		vals.Set("inactiveReasons", p.InactiveReasons)
	}
	if p.IncludeThirdPartyClaims {
		vals.Set("includeThirdPartyClaims", "true")
	}
	if p.IsVideoShortsEligible {
		vals.Set("isVideoShortsEligible", "true")
	}
	if p.OnBehalfOfContentOwner != "" {
		vals.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.Origin != "" {
		vals.Set("origin", p.Origin)
	}
	if p.PageToken != "" {
		vals.Set("pageToken", p.PageToken)
	}
	if p.PartnerUploaded {
		vals.Set("partnerUploaded", "true")
	}
	if p.Q != "" {
		vals.Set("q", p.Q)
	}
	if p.ReferenceId != "" {
		vals.Set("referenceId", p.ReferenceId)
	}
	if p.Sort != "" {
		vals.Set("sort", p.Sort)
	}
	if p.Status != "" {
		vals.Set("status", string(p.Status))
	}
	if p.StatusModifiedAfter != "" && DateRegexp.MatchString(p.StatusModifiedAfter) {
		vals.Set("statusModifiedAfter", p.StatusModifiedAfter)
	}
	if p.StatusModifiedBefore != "" && DateRegexp.MatchString(p.StatusModifiedBefore) {
		vals.Set("statusModifiedBefore", p.StatusModifiedBefore)
	}
	if len(p.VideoIds) > 0 {
		vals.Set("videoId", strings.Join(p.VideoIds, ","))
	}
	return vals
}

// SearchClaims retrieves a list of claims that match the search criteria. The
// API response uses pagination. You must specify one and only one search filter:
// assetId, videoId, q, referenceId, or status.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claimSearch/list
func SearchClaims(runner RequestRunner, p *SearchClaimsParams) (*ClaimSearchResponse, error) {
	if !p.Validate() {
		return nil, ErrInvalidClaimSearchParams
	}
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    SearchClaimsUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ClaimSearchResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── claims.get ──────────────────────────────────────────────────────────────

// GetClaimParams are parameters for the claims.get method, which retrieves a
// specific claim by ID.
type GetClaimParams struct {
	// ClaimId: The claimId parameter specifies the claim ID of the claim being
	// retrieved. This is a required parameter.
	ClaimId string

	// OnBehalfOfContentOwner: The onBehalfOfContentOwner parameter identifies
	// the content owner that the user is acting on behalf of. This parameter
	// supports users whose accounts are associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetClaimParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetClaim retrieves a specific claim by ID. The API response contains the
// claim's current status and other details.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claims/get
func GetClaim(runner RequestRunner, p *GetClaimParams) (*Claim, error) {
	if p.ClaimId == "" {
		return nil, ErrMissingClaimId
	}
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ClaimsUrl + "/" + p.ClaimId,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out Claim
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── claims.list ─────────────────────────────────────────────────────────────

// ListClaimsParams are parameters for the claims.list method, which retrieves a
// list of claims administered by the content owner associated with the
// currently authenticated user.
type ListClaimsParams struct {
	// AssetId: Use the claimSearch.list method's assetId parameter to search for
	// claim snippets by asset ID. Use this method's assetId parameter to
	// retrieve the actual claim resources for a specific asset.
	AssetId string

	// Id: The id parameter specifies a list of comma-separated YouTube claim
	// IDs to retrieve.
	Id string

	// OnBehalfOfContentOwner: The onBehalfOfContentOwner parameter identifies
	// the content owner that the user is acting on behalf of. This parameter
	// supports users whose accounts are associated with multiple content owners.
	OnBehalfOfContentOwner string

	// PageToken: The pageToken parameter specifies a token that identifies a
	// particular page of results to return. For example, set this parameter to
	// the value of the nextPageToken from the previous API response to retrieve
	// the next page of results.
	PageToken string

	// Q: Use the claimSearch.list method's q parameter to search for claim
	// snippets that match a particular query string. Use this method's q
	// parameter to retrieve the actual claim resources that are associated with
	// a specified asset.
	Q string

	// VideoId: The videoId parameter specifies the YouTube video ID of the video
	// for which you are retrieving claims.
	VideoId string
}

func (p *ListClaimsParams) Values() url.Values {
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
	if p.Q != "" {
		v.Set("q", p.Q)
	}
	if p.VideoId != "" {
		v.Set("videoId", p.VideoId)
	}
	return v
}

// ListClaims retrieves a list of claims administered by the content owner
// associated with the currently authenticated user. The API response uses
// pagination.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claims/list
func ListClaims(runner RequestRunner, p *ListClaimsParams) (*ClaimListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ClaimsUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ClaimListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── claims.insert ───────────────────────────────────────────────────────────

// InsertClaimParams are parameters for the claims.insert method, which creates
// a new claim. The authenticated user must be associated with a content owner.
type InsertClaimParams struct {
	// OnBehalfOfContentOwner: The onBehalfOfContentOwner parameter identifies
	// the content owner that the user is acting on behalf of. This parameter
	// supports users whose accounts are associated with multiple content owners.
	OnBehalfOfContentOwner string

	// IsManualClaim: The isManualClaim parameter indicates whether the claim is
	// a manual claim. Manual claims are claims that are not initiated through
	// Content ID matching.
	IsManualClaim bool

	// Claim: The claim resource to insert.
	Claim *Claim
}

func (p *InsertClaimParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.IsManualClaim {
		v.Set("isManualClaim", "true")
	}
	return v
}

// InsertClaim creates a new claim. You can create a claim for a video that has
// already been uploaded to YouTube. The request body contains a claim resource
// in which you specify the asset, video, content type, and policy.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claims/insert
func InsertClaim(runner RequestRunner, p *InsertClaimParams) (*Claim, error) {
	body, err := jsonBody(p.Claim)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    ClaimsUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Claim
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── claims.patch ────────────────────────────────────────────────────────────

// PatchClaimsParams are parameters for the claims.patch method, which patches
// an existing claim by updating the fields included in the request body. This
// method supports patch semantics.
type PatchClaimsParams struct {
	// ClaimId: The claimId parameter specifies the claim ID of the claim being
	// patched. This is a required parameter.
	ClaimId string

	// OnBehalfOfContentOwner: The onBehalfOfContentOwner parameter identifies
	// the content owner that the user is acting on behalf of. This parameter
	// supports users whose accounts are associated with multiple content owners.
	OnBehalfOfContentOwner string

	// Status: The status to set on the claim. The claim's status must be valid
	// per the ClaimStatus enumeration.
	Status ClaimStatus
}

func (p *PatchClaimsParams) Validate() bool {
	return p.ClaimId != "" && p.Status != "" && p.Status.Valid()
}

func (p *PatchClaimsParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

func (p *PatchClaimsParams) Body() (io.Reader, error) {
	m := make(map[string]interface{})
	if p.Status != "" && p.Status.Valid() {
		m["status"] = p.Status
	}
	if len(m) == 0 {
		return nil, ErrInvalidPatchClaimsParams
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// PatchClaim patches an existing claim by only updating the fields included in
// the request body. Use this method to update a claim's status. This method
// supports patch semantics -- only the fields included in the request body are
// updated.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claims/patch
func PatchClaim(runner RequestRunner, p *PatchClaimsParams) (*Claim, error) {
	if !p.Validate() {
		return nil, ErrInvalidPatchClaimsParams
	}
	body, err := p.Body()
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPatch,
		Url:    ClaimsUrl + "/" + p.ClaimId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Claim
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── claims.update ───────────────────────────────────────────────────────────

// UpdateClaimParams are parameters for the claims.update method, which updates
// an existing claim. The request replaces the existing claim resource entirely,
// so all fields must be set in the request body.
type UpdateClaimParams struct {
	// ClaimId: The claimId parameter specifies the claim ID of the claim being
	// updated. This is a required parameter.
	ClaimId string

	// OnBehalfOfContentOwner: The onBehalfOfContentOwner parameter identifies
	// the content owner that the user is acting on behalf of. This parameter
	// supports users whose accounts are associated with multiple content owners.
	OnBehalfOfContentOwner string

	// Claim: The claim resource with updated fields. This replaces the existing
	// resource entirely.
	Claim *Claim
}

func (p *UpdateClaimParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// UpdateClaim updates an existing claim by replacing the claim resource
// entirely. You must set all fields in the request body. If you only need to
// update specific fields, use PatchClaim instead.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claims/update
func UpdateClaim(runner RequestRunner, p *UpdateClaimParams) (*Claim, error) {
	if p.ClaimId == "" {
		return nil, ErrMissingClaimId
	}
	body, err := jsonBody(p.Claim)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPut,
		Url:    ClaimsUrl + "/" + p.ClaimId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Claim
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── claimHistory.get ────────────────────────────────────────────────────────

// GetClaimHistoryParams are parameters for the claimHistory.get method, which
// retrieves the claim history for a specific claim.
type GetClaimHistoryParams struct {
	// ClaimId: The claimId parameter specifies the YouTube claim ID of the claim
	// for which you are retrieving the claim history. This is a required
	// parameter.
	ClaimId string

	// OnBehalfOfContentOwner: The onBehalfOfContentOwner parameter identifies
	// the content owner that the user is acting on behalf of. This parameter
	// supports users whose accounts are associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetClaimHistoryParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetClaimHistory retrieves the claim history for a specific claim. The claim
// history lists the actions taken on a claim over time, including status changes,
// policy changes, and dispute events.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/claimHistory/get
func GetClaimHistory(runner RequestRunner, p *GetClaimHistoryParams) (*ClaimHistory, error) {
	if p.ClaimId == "" {
		return nil, ErrMissingClaimId
	}
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ClaimHistoryUrl + "/" + p.ClaimId,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ClaimHistory
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── helpers ─────────────────────────────────────────────────────────────────

// jsonBody marshals v to JSON and returns a reader.
func jsonBody(v interface{}) (io.Reader, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
