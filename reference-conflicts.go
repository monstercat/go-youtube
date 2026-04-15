package youtube

import (
	"net/http"
	"net/url"
)

const (
	ReferenceConflictsUrl = YoutubePartnerV1 + "/referenceConflicts"
)

// ReferenceConflict represents a conflict between two references that identify
// the same content. When YouTube identifies a conflict, it blocks the newer
// reference until the conflict is resolved.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/referenceConflicts#ReferenceConflict
type ReferenceConflict struct {
	// ConflictingReferenceId is the ID of the newer reference that caused the conflict.
	ConflictingReferenceId string `json:"conflictingReferenceId,omitempty"`
	// ExpiryTime is the date and time when the conflict will expire and the
	// newer reference will be automatically activated if the conflict is not resolved.
	ExpiryTime string `json:"expiryTime,omitempty"`
	// Id is the unique ID that YouTube assigns to the reference conflict.
	Id string `json:"id,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#referenceConflict".
	Kind string `json:"kind,omitempty"`
	// Matches is a list of match segments between the conflicting and original references.
	Matches []*ReferenceConflictMatch `json:"matches,omitempty"`
	// OriginalReferenceId is the ID of the existing (original) reference that
	// the newer reference conflicts with.
	OriginalReferenceId string `json:"originalReferenceId,omitempty"`
	// Status is the conflict's status. Valid values are "conflicting" and "resolved".
	Status string `json:"status,omitempty"`
}

// ReferenceConflictMatch describes a matched segment between two conflicting
// references, including offset and length information for each reference.
type ReferenceConflictMatch struct {
	// ConflictingReferenceOffsetMs is the offset in milliseconds into the
	// conflicting reference at which the matched segment begins.
	ConflictingReferenceOffsetMs string `json:"conflicting_reference_offset_ms,omitempty"`
	// LengthMs is the length of the conflicting match segment in milliseconds.
	LengthMs string `json:"length_ms,omitempty"`
	// OriginalReferenceOffsetMs is the offset in milliseconds into the original
	// reference at which the matched segment begins.
	OriginalReferenceOffsetMs string `json:"original_reference_offset_ms,omitempty"`
	// Type is the type of match (audio or video).
	Type string `json:"type,omitempty"`
}

// ReferenceConflictListResponse is the response for referenceConflicts.list.
type ReferenceConflictListResponse struct {
	// Items is a list of reference conflicts that match the request criteria.
	Items []*ReferenceConflict `json:"items"`
	// Kind identifies the API resource type. The value is "youtubePartner#referenceConflictList".
	Kind string `json:"kind,omitempty"`
	// NextPageToken is the token that can be used as the value of the pageToken
	// parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`
	// PageInfo contains paging details for the result set.
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
}

// ── referenceConflicts.get ──────────────────────────────────────────────────

// GetReferenceConflictParams are parameters for referenceConflicts.get.
type GetReferenceConflictParams struct {
	// ReferenceConflictId is the ID of the reference conflict to retrieve.
	ReferenceConflictId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetReferenceConflictParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetReferenceConflict retrieves information about the specified reference
// conflict. A reference conflict occurs when a new reference identifies content
// that already matches an existing reference.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/referenceConflicts/get
func GetReferenceConflict(runner RequestRunner, p *GetReferenceConflictParams) (*ReferenceConflict, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ReferenceConflictsUrl + "/" + p.ReferenceConflictId,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ReferenceConflict
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── referenceConflicts.list ─────────────────────────────────────────────────

// ListReferenceConflictsParams are parameters for referenceConflicts.list.
type ListReferenceConflictsParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// PageToken is the token that identifies a specific page in the result set
	// that should be returned. Set this to the value of nextPageToken from a
	// previous response to retrieve the next page.
	PageToken string
}

func (p *ListReferenceConflictsParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.PageToken != "" {
		v.Set("pageToken", p.PageToken)
	}
	return v
}

// ListReferenceConflicts retrieves a list of unresolved reference conflicts.
// Conflicts are returned in order of when the conflicting reference was created,
// starting with the most recently created.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/referenceConflicts/list
func ListReferenceConflicts(runner RequestRunner, p *ListReferenceConflictsParams) (*ReferenceConflictListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ReferenceConflictsUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ReferenceConflictListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
