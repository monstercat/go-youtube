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
)

// ClaimStatus - status of a claim. Acceptable values are:
// - active – Restrict results to claims with active status.
//  - appealed – Restrict results to claims with appealed status.
// - disputed – Restrict results to claims with disputed status.
//  - inactive – Restrict results to claims with inactive status.
//  - pending – Restrict results to claims with pending status.
// - potential – Restrict results to claims with potetial status.
// - routedForReview – Restrict results to claims that require review based on a match policy rule.
// - takedown – Restrict results to claims with takedown status.
type ClaimStatus string

const (
	SearchClaimsUrl = YoutubePartnerV1 + "/claimSearch"
	PatchClaimUrl   = YoutubePartnerV1 + "/claims"

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
	DateRegexp = regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}")
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

type Claim struct {
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

	// Status: The claim's status. When updating a claim, you can update its
	// status from active to inactive to effectively release the claim, but
	// the API does not support other updates to a claim's status.
	Status ClaimStatus `json:"status,omitempty"`

	// TimeCreated: The time the claim was created.
	TimeCreated string `json:"timeCreated,omitempty"`

	// TimeStatusLastModified: The time that the claim's status was last modified.
	TimeStatusLastModified string `json:"timeStatusLastModified"`

	// VideoId: The unique YouTube video ID that identifies the video
	// associated with the claim.
	VideoId string `json:"videoId,omitempty"`
}

type SearchClaimsResponse struct {
	// Items are the returned claims
	Items []*Claim `json:"items"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken"`

	// PageInfo: General pagination information.
	PageInfo *PageInfo `json:"pageInfo"`
}

// SearchClaimsParams are parameters for the search claims method.
// see https://developers.google.com/youtube/partner/docs/v1/claimSearch/list
//
// Filter parameters for searching by ID or query string
// (Specify exactly one of the following parameters. Don't use these parameters to search by date or status.)
// - assetId
// - q
// - referenceId
// - videoId
//
// Filter parameters for searching by date or status
// (Specify at least one of the following parameters. These can also be used as optional parameters to search by ID or query string.)
// - status
type SearchClaimsParams struct {
	// The AssetId parameter specifies the YouTube asset ID of the asset for which you are retrieving claims.
	//
	// Content partners who are managing composition assets can refer to the Managing composition assets guide to
	// understand how the API response content differs if the parameter value identifies a composition view or a
	//composition share.
	AssetId string

	// The Q parameter specifies the query string to use to filter search results. YouTube searches for the query
	// string in the following claim fields: video_title, video_keywords, user_name, isrc, iswc, grid, custom_id,
	// and in the content owner's email address.
	Q string

	// The ReferenceId parameter specifies the YouTube reference ID of the reference for which you are retrieving claims.
	ReferenceId string

	// The VideoIds parameter specifies a comma-separated list of up to 10 YouTube video IDs for which you are
	// retrieving claims.
	VideoIds []string

	// The PageToken parameter specifies a token that identifies a particular page of results to return. For example,
	// set this parameter to the value of the nextPageToken value from the previous API response to retrieve the
	// next page of search results.
	PageToken string

	// The Status parameter restricts your results to only claims in the specified status.
	//
	// Acceptable values are:
	// - active – Restrict results to claims with active status.
	// - appealed – Restrict results to claims with appealed status.
	// - disputed – Restrict results to claims with disputed status.
	// - inactive – Restrict results to claims with inactive status.
	// - pending – Restrict results to claims with pending status.
	// - potential – Restrict results to claims with potetial status.
	// - routedForReview – Restrict results to claims that require review based on a match policy rule.
	// - takedown – Restrict results to claims with takedown status.
	Status ClaimStatus

	// Used along with the videoId parameter this parameter determines whether to include third party claims
	// in the search results.
	//
	// This parameter is only supported for searching by ID or query string.
	IncludeThirdPartyClaims bool

	// The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This
	// parameter supports users whose accounts are associated with multiple content owners.
	OnBehalfOfContentOwner string

	// The sort parameter specifies the method that will be used to order resources in the API response. The default
	// value is date. However, if the status parameter value is either appealed, disputed, pending, potential, or
	// routedForReview, then results will be sorted by the time that the claim review period expires.
	Sort string

	// The statusModifiedAfter parameter allows you to restrict the result set to only include claims that have had
	// their status modified on or after the specified date (inclusive). The date specified must be on or after
	// June 30, 2016 (2016-06-30). The parameter value's format is YYYY-MM-DD.
	StatusModifiedAfter string
}

func (p *SearchClaimsParams) Validate() bool {
	// Exactly one of the following is required:
	// - assetId
	// - q
	// - referenceId
	// - videoId
	//
	// If status is provided, it is not necessary to have one of the above.
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

	// At least one of the following is required:
	// - status
	if p.Status == "" || !p.Status.Valid() {
		return false
	}

	return true
}

func (p *SearchClaimsParams) Values() url.Values {
	vals := url.Values{}
	if p.AssetId != "" {
		vals.Add("assetId", p.AssetId)
	}
	if p.Q != "" {
		vals.Add("q", p.Q)
	}
	if p.ReferenceId != "" {
		vals.Add("referenceId", p.ReferenceId)
	}
	if len(p.VideoIds) > 0 {
		vals.Add("videoId", strings.Join(p.VideoIds, ","))
	}
	if p.Status != "" {
		vals.Add("status", string(p.Status))
	}
	if p.IncludeThirdPartyClaims {
		vals.Add("includeThirdPartyClaims", "true")
	} else {
		vals.Add("includeThirdPartyClaims", "false")
	}
	if p.OnBehalfOfContentOwner != "" {
		vals.Add("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.PageToken != "" {
		vals.Add("pageToken", p.PageToken)
	}
	switch p.Sort {
	case ClaimSortDate:
		vals.Add("sort", "date")
	case ClaimSortViewCount:
		vals.Add("sort", "viewCount")
	}
	if p.StatusModifiedAfter != "" && DateRegexp.MatchString(p.StatusModifiedAfter) {
		vals.Add("statusModifiedAfter", p.StatusModifiedAfter)
	}
	return vals
}

// SearchClaims - Retrieves a list of claims that match the search criteria. There are two main ways you can use this
// method to search for claims:
// - Search for claims that are associated with a specific asset, reference or video, or that match a specified query string.
// - Search for claims that were created or modified within a given timeframe, or that have a specific status.
//
// Authorization
// =============
// This request requires authorization with at least one of the following scopes:
// - https://www.googleapis.com/auth/youtubepartner
//
// see https://developers.google.com/youtube/partner/docs/v1/claimSearch/list
func SearchClaims(runner RequestRunner, p *SearchClaimsParams) (*SearchClaimsResponse, error) {
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

	var out SearchClaimsResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// PatchClaimsParams are params related to the patch claims call.
type PatchClaimsParams struct {
	// The ClaimId parameter specifies the claim ID of the claim being updated.
	// REQUIRED
	ClaimId string

	// The OnBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of. This
	// parameter supports users whose accounts are associated with multiple content owners.
	OnBehalfOfContentOwner string

	// The ClaimStatus to set (optional).
	// The claim's status. When updating a claim, you can update its status from active to inactive to effectively
	// release the claim.
	Status ClaimStatus

	// TODO: policy
	// TODO: blockOutsideOwnership
}

func (p *PatchClaimsParams) Validate() bool {
	return p.ClaimId != "" && p.Status != "" && p.Status.Valid()
}

func (p *PatchClaimsParams) Url() string {
	return PatchClaimUrl + "/" + p.ClaimId
}

func (p *PatchClaimsParams) Params() url.Values {
	v := make(url.Values)
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

// PatchClaims - Updates an existing claim by either changing its policy or its status. You can update a claim's status
// from active to inactive to effectively release the claim. This method supports patch semantics.
//
// Authorization
// =============
// This request requires authorization with at least one of the following scopes:
// - https://www.googleapis.com/auth/youtubepartner
//
// https://developers.google.com/youtube/partner/docs/v1/claims/patch
func PatchClaims(runner RequestRunner, p *PatchClaimsParams) (*Claim, error) {
	if !p.Validate() {
		return nil, ErrInvalidPatchClaimsParams
	}
	body, err := p.Body()
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPatch,
		Url:    p.Url(),
		Params: p.Params(),
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
