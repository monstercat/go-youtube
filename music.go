package youtube

import (
	"net/http"
	"net/url"
	"strconv"
)

const (
	MusicChangeRequestsUrl = YoutubePartnerV1 + "/music/changeRequests"
	MusicReleasesUrl       = YoutubePartnerV1 + "/music/releases"
)

// Artist represents a music artist.
type Artist struct {
	Name string `json:"name,omitempty"`
}

// MusicTrack represents a music track in the YouTube Music system. A track
// is associated with a specific release and may have an external video ID
// linking it to a YouTube video.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/musicTracks#MusicTrack
type MusicTrack struct {
	// Artists is the list of artists who performed the track.
	Artists []*Artist `json:"artists,omitempty"`
	// ExternalVideoId is the YouTube video ID associated with the track, if any.
	ExternalVideoId string `json:"externalVideoId,omitempty"`
	// HasOpenChangeRequest indicates whether the track has a pending (open) change request.
	HasOpenChangeRequest bool `json:"hasOpenChangeRequest,omitempty"`
	// Name is the resource name of the track.
	Name string `json:"name,omitempty"`
	// Title is the title of the track.
	Title string `json:"title,omitempty"`
}

// ListMusicTracksResponse is the response for musicTracks.list.
type ListMusicTracksResponse struct {
	// NextPageToken is the token that can be used as the value of the pageToken
	// parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`
	// Tracks is the list of music tracks that match the request criteria.
	Tracks []*MusicTrack `json:"tracks,omitempty"`
}

// MusicRelease represents a music release (album, single, EP, etc.) in the
// YouTube Music system. Each release may contain multiple tracks and is
// associated with a YouTube playlist.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/musicReleases#MusicRelease
type MusicRelease struct {
	// Artists is the list of artists associated with the release.
	Artists []*Artist `json:"artists,omitempty"`
	// HasOpenChangeRequest indicates whether the release has a pending (open) change request.
	HasOpenChangeRequest bool `json:"hasOpenChangeRequest,omitempty"`
	// Name is the resource name of the release.
	Name string `json:"name,omitempty"`
	// PlaylistId is the YouTube playlist ID for this release, if available.
	PlaylistId string `json:"playlistId,omitempty"`
	// Title is the title of the release.
	Title string `json:"title,omitempty"`
}

// ListMusicReleasesResponse is the response for musicReleases.list.
type ListMusicReleasesResponse struct {
	// NextPageToken is the token that can be used as the value of the pageToken
	// parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`
	// Releases is the list of music releases that match the request criteria.
	Releases []*MusicRelease `json:"releases,omitempty"`
}

// MusicChangeRequest represents a request to change music metadata, correct
// artist reconciliation, fix playability issues, or address other music data
// problems. The type field indicates which kind of change is requested.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/musicChangeRequests#MusicChangeRequest
type MusicChangeRequest struct {
	// CreateTime is the date and time the change request was created.
	CreateTime string `json:"createTime,omitempty"`
	// IncorrectMetadata provides the corrected metadata when the change request
	// type is for incorrect metadata (title or artist corrections).
	IncorrectMetadata *IncorrectMetadata `json:"incorrectMetadata,omitempty"`
	// IncorrectMusicVideo provides the desired video ID when the change request
	// type is for an incorrect music video association.
	IncorrectMusicVideo *DesiredMusicVideo `json:"incorrectMusicVideo,omitempty"`
	// IncorrectPlayability provides an explanation when the change request type
	// is for incorrect playability settings.
	IncorrectPlayability *IncorrectPlayability `json:"incorrectPlayability,omitempty"`
	// MisreconciledArtist provides the desired artist name when the change
	// request type is for a misreconciled artist page.
	MisreconciledArtist *DesiredArtist `json:"misreconciledArtist,omitempty"`
	// Name is the resource name of the change request.
	Name string `json:"name,omitempty"`
	// Release is the resource name of the release associated with the change request.
	Release string `json:"release,omitempty"`
	// State is the current state of the change request (e.g., "open", "approved", "rejected").
	State string `json:"state,omitempty"`
	// Track is the resource name of the track associated with the change request.
	Track string `json:"track,omitempty"`
	// Type is the type of change request (e.g., "incorrectMetadata",
	// "incorrectMusicVideo", "incorrectPlayability", "misreconciledArtist",
	// "undesiredDiscography").
	Type string `json:"type,omitempty"`
	// UndesiredDiscography provides an explanation when the change request type
	// is for removing undesired content from an artist's discography.
	UndesiredDiscography *UndesiredDiscography `json:"undesiredDiscography,omitempty"`
}

// IncorrectMetadata describes incorrect metadata in a music change request
// and provides the corrected values.
type IncorrectMetadata struct {
	// CorrectTitle is the corrected title for the track.
	CorrectTitle string `json:"correctTitle,omitempty"`
	// CorrectArtist is the corrected list of artists for the track.
	CorrectArtist []*Artist `json:"correctArtist,omitempty"`
}

// DesiredMusicVideo describes the desired music video in a change request.
type DesiredMusicVideo struct {
	// DesiredVideoId is the YouTube video ID that should be associated with the track.
	DesiredVideoId string `json:"desiredVideoId,omitempty"`
}

// IncorrectPlayability describes incorrect playability in a change request.
type IncorrectPlayability struct {
	// Explanation is a free-text explanation of the playability issue.
	Explanation string `json:"explanation,omitempty"`
}

// DesiredArtist describes a desired artist correction in a change request.
type DesiredArtist struct {
	// DesiredArtistName is the correct artist name that the track should be reconciled to.
	DesiredArtistName string `json:"desiredArtistName,omitempty"`
}

// UndesiredDiscography describes an undesired discography in a change request.
type UndesiredDiscography struct {
	// Explanation is a free-text explanation of why the content should be removed
	// from the artist's discography.
	Explanation string `json:"explanation,omitempty"`
}

// ListMusicChangeRequestsResponse is the response for musicChangeRequests.list.
type ListMusicChangeRequestsResponse struct {
	// ChangeRequests is the list of music change requests that match the request criteria.
	ChangeRequests []*MusicChangeRequest `json:"changeRequests,omitempty"`
	// NextPageToken is the token that can be used as the value of the pageToken
	// parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// ── musicTracks.list ────────────────────────────────────────────────────────

// ListMusicTracksParams are parameters for musicTracks.list.
type ListMusicTracksParams struct {
	// Parent is the resource name of the parent release whose tracks are being listed.
	Parent string
	// FilterArtistNameMatches filters for tracks whose artist name matches the specified string.
	FilterArtistNameMatches string
	// FilterExternalVideoIds filters for tracks with the specified external
	// (YouTube) video IDs. Comma-separated list.
	FilterExternalVideoIds string
	// FilterHasClosedChangeReq filters for tracks that have a closed change request.
	FilterHasClosedChangeReq bool
	// FilterHasOpenChangeReq filters for tracks that have an open change request.
	FilterHasOpenChangeReq bool
	// FilterIsrcs filters for tracks with the specified ISRC codes. Comma-separated list.
	FilterIsrcs string
	// FilterTitleMatches filters for tracks whose title matches the specified string.
	FilterTitleMatches string
	// FilterUpcs filters for tracks with the specified UPC codes. Comma-separated list.
	FilterUpcs string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// PageSize is the maximum number of items that should be returned in the result set.
	PageSize int
	// PageToken is the token that identifies a specific page in the result set
	// that should be returned.
	PageToken string
}

func (p *ListMusicTracksParams) Values() url.Values {
	v := url.Values{}
	if p.FilterArtistNameMatches != "" {
		v.Set("filter.artistNameMatches", p.FilterArtistNameMatches)
	}
	if p.FilterExternalVideoIds != "" {
		v.Set("filter.externalVideoIds", p.FilterExternalVideoIds)
	}
	if p.FilterHasClosedChangeReq {
		v.Set("filter.hasClosedChangeRequest", "true")
	}
	if p.FilterHasOpenChangeReq {
		v.Set("filter.hasOpenChangeRequest", "true")
	}
	if p.FilterIsrcs != "" {
		v.Set("filter.isrcs", p.FilterIsrcs)
	}
	if p.FilterTitleMatches != "" {
		v.Set("filter.titleMatches", p.FilterTitleMatches)
	}
	if p.FilterUpcs != "" {
		v.Set("filter.upcs", p.FilterUpcs)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.PageSize > 0 {
		v.Set("pageSize", strconv.Itoa(p.PageSize))
	}
	if p.PageToken != "" {
		v.Set("pageToken", p.PageToken)
	}
	return v
}

func musicTracksUrl(parent string) string {
	return YoutubePartnerV1 + "/music/" + parent + "/tracks"
}

// ListMusicTracks retrieves a list of music tracks owned by the content owner
// and associated with a given release. Results can be filtered by artist name,
// title, ISRC, UPC, or change request status.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/musicTracks/list
func ListMusicTracks(runner RequestRunner, p *ListMusicTracksParams) (*ListMusicTracksResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    musicTracksUrl(p.Parent),
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ListMusicTracksResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── musicReleases.list ──────────────────────────────────────────────────────

// ListMusicReleasesParams are parameters for musicReleases.list.
type ListMusicReleasesParams struct {
	// FilterArtistNameMatches filters for releases whose artist name matches the specified string.
	FilterArtistNameMatches string
	// FilterHasClosedChangeReq filters for releases that have a closed change request.
	FilterHasClosedChangeReq bool
	// FilterHasOpenChangeReq filters for releases that have an open change request.
	FilterHasOpenChangeReq bool
	// FilterTitleMatches filters for releases whose title matches the specified string.
	FilterTitleMatches string
	// FilterUpcs filters for releases with the specified UPC codes. Comma-separated list.
	FilterUpcs string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// PageSize is the maximum number of items that should be returned in the result set.
	PageSize int
	// PageToken is the token that identifies a specific page in the result set
	// that should be returned.
	PageToken string
}

func (p *ListMusicReleasesParams) Values() url.Values {
	v := url.Values{}
	if p.FilterArtistNameMatches != "" {
		v.Set("filter.artistNameMatches", p.FilterArtistNameMatches)
	}
	if p.FilterHasClosedChangeReq {
		v.Set("filter.hasClosedChangeRequest", "true")
	}
	if p.FilterHasOpenChangeReq {
		v.Set("filter.hasOpenChangeRequest", "true")
	}
	if p.FilterTitleMatches != "" {
		v.Set("filter.titleMatches", p.FilterTitleMatches)
	}
	if p.FilterUpcs != "" {
		v.Set("filter.upcs", p.FilterUpcs)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.PageSize > 0 {
		v.Set("pageSize", strconv.Itoa(p.PageSize))
	}
	if p.PageToken != "" {
		v.Set("pageToken", p.PageToken)
	}
	return v
}

// ListMusicReleases retrieves a list of music releases owned by the content
// owner. Results can be filtered by artist name, title, UPC, or change
// request status.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/musicReleases/list
func ListMusicReleases(runner RequestRunner, p *ListMusicReleasesParams) (*ListMusicReleasesResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    MusicReleasesUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ListMusicReleasesResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── musicChangeRequests.list ────────────────────────────────────────────────

// ListMusicChangeRequestsParams are parameters for musicChangeRequests.list.
type ListMusicChangeRequestsParams struct {
	// FilterParent filters for change requests associated with the specified
	// parent resource name (release or track).
	FilterParent string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// PageSize is the maximum number of items that should be returned in the result set.
	PageSize int
	// PageToken is the token that identifies a specific page in the result set
	// that should be returned.
	PageToken string
}

func (p *ListMusicChangeRequestsParams) Values() url.Values {
	v := url.Values{}
	if p.FilterParent != "" {
		v.Set("filter.parent", p.FilterParent)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.PageSize > 0 {
		v.Set("pageSize", strconv.Itoa(p.PageSize))
	}
	if p.PageToken != "" {
		v.Set("pageToken", p.PageToken)
	}
	return v
}

// ListMusicChangeRequests retrieves a list of music change requests for the
// content owner. Results can be filtered by parent resource (release or track).
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/musicChangeRequests/list
func ListMusicChangeRequests(runner RequestRunner, p *ListMusicChangeRequestsParams) (*ListMusicChangeRequestsResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    MusicChangeRequestsUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ListMusicChangeRequestsResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── musicChangeRequests.create ──────────────────────────────────────────────

// CreateMusicChangeRequestParams are parameters for musicChangeRequests.create.
type CreateMusicChangeRequestParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// ChangeRequest is the MusicChangeRequest resource to create.
	ChangeRequest *MusicChangeRequest
}

func (p *CreateMusicChangeRequestParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// CreateMusicChangeRequest creates a new music change request. Change requests
// can be used to correct metadata (title/artist), fix artist reconciliation,
// update music video associations, fix playability issues, or remove undesired
// content from an artist's discography.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/musicChangeRequests/create
func CreateMusicChangeRequest(runner RequestRunner, p *CreateMusicChangeRequestParams) (*MusicChangeRequest, error) {
	body, err := jsonBody(p.ChangeRequest)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    MusicChangeRequestsUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out MusicChangeRequest
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
