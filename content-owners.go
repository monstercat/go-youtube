package youtube

import (
	"net/http"
	"net/url"
)

const (
	ContentOwnersUrl = YoutubePartnerV1 + "/contentOwners"
)

// ContentOwner represents a YouTube content owner. A content owner is the
// entity that has the right to monetize, track, or block videos on YouTube.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/contentOwners#ContentOwner
type ContentOwner struct {
	// ConflictNotificationEmail is the email address to which YouTube sends
	// notifications about reference conflicts.
	ConflictNotificationEmail string `json:"conflictNotificationEmail,omitempty"`
	// DisplayName is the content owner's display name.
	DisplayName string `json:"displayName,omitempty"`
	// DisputeNotificationEmails is the list of email addresses to which YouTube
	// sends notifications when a partner disputes a claim.
	DisputeNotificationEmails []string `json:"disputeNotificationEmails,omitempty"`
	// FingerprintReportNotificationEmails is the list of email addresses to
	// which YouTube sends fingerprint reports.
	FingerprintReportNotificationEmails []string `json:"fingerprintReportNotificationEmails,omitempty"`
	// Id is the unique ID that YouTube assigns to the content owner.
	Id string `json:"id,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#contentOwner".
	Kind string `json:"kind,omitempty"`
	// PrimaryNotificationEmails is the list of email addresses to which YouTube
	// sends primary notifications about the content owner's account activity.
	PrimaryNotificationEmails []string `json:"primaryNotificationEmails,omitempty"`
}

// ContentOwnerListResponse is the response for contentOwners.list.
type ContentOwnerListResponse struct {
	// Items is a list of content owners that match the request criteria.
	Items []*ContentOwner `json:"items"`
	// Kind identifies the API resource type. The value is "youtubePartner#contentOwnerList".
	Kind string `json:"kind,omitempty"`
}

// ── contentOwners.get ───────────────────────────────────────────────────────

// GetContentOwnerParams are parameters for contentOwners.get.
type GetContentOwnerParams struct {
	// ContentOwnerId is the unique ID of the content owner to retrieve.
	ContentOwnerId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetContentOwnerParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetContentOwner retrieves information about the specified content owner,
// including display name, notification email addresses, and conflict
// notification settings.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/contentOwners/get
func GetContentOwner(runner RequestRunner, p *GetContentOwnerParams) (*ContentOwner, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ContentOwnersUrl + "/" + p.ContentOwnerId,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ContentOwner
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── contentOwners.list ──────────────────────────────────────────────────────

// ListContentOwnersParams are parameters for contentOwners.list.
type ListContentOwnersParams struct {
	// FetchMine indicates whether to retrieve the content owner associated with
	// the currently authenticated user. Set to true to retrieve content owners
	// that the authenticated user is able to act on behalf of.
	FetchMine bool
}

func (p *ListContentOwnersParams) Values() url.Values {
	v := url.Values{}
	if p.FetchMine {
		v.Set("fetchMine", "true")
	}
	return v
}

// ListContentOwners retrieves a list of content owners that match the request
// criteria. Most commonly used with FetchMine=true to retrieve the content
// owners associated with the currently authenticated user.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/contentOwners/list
func ListContentOwners(runner RequestRunner, p *ListContentOwnersParams) (*ContentOwnerListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ContentOwnersUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ContentOwnerListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
