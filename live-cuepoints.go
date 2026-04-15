package youtube

import (
	"net/http"
	"net/url"
)

const (
	LiveCuepointsUrl = YoutubePartnerV1 + "/liveCuepoints"
)

// LiveCuepoint represents a cuepoint that can be inserted into a live broadcast.
// Cuepoints trigger ad breaks during a live stream.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/liveCuepoints#LiveCuepoint
type LiveCuepoint struct {
	// BroadcastId is the ID of the live broadcast into which the cuepoint is inserted.
	BroadcastId string `json:"broadcastId,omitempty"`
	// Id is the unique ID that YouTube assigns to the cuepoint.
	Id string `json:"id,omitempty"`
	// Kind identifies the API resource type. The value is "youtubePartner#liveCuepoint".
	Kind string `json:"kind,omitempty"`
	// Settings contains the cuepoint configuration, including type, duration, and timing.
	Settings *CuepointSettings `json:"settings,omitempty"`
}

// CuepointSettings contains the settings for a live cuepoint, including type,
// duration, and timing.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/liveCuepoints#CuepointSettings
type CuepointSettings struct {
	// CueType is the type of cuepoint. Valid values are "ad" (an ad break).
	CueType string `json:"cueType,omitempty"`
	// DurationSecs is the duration of the ad break in seconds. If not
	// specified, the ad break is open-ended.
	DurationSecs int `json:"durationSecs,omitempty"`
	// OffsetTimeMs is the offset from the start of the broadcast, in
	// milliseconds, at which the cuepoint should be inserted.
	OffsetTimeMs string `json:"offsetTimeMs,omitempty"`
	// Walltime is the wall clock time at which the cuepoint should be inserted.
	// This is an RFC 3339 timestamp.
	Walltime string `json:"walltime,omitempty"`
}

// ── liveCuepoints.insert ────────────────────────────────────────────────────

// InsertLiveCuepointParams are parameters for liveCuepoints.insert.
type InsertLiveCuepointParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// Cuepoint is the LiveCuepoint resource to insert into the broadcast.
	Cuepoint *LiveCuepoint
}

func (p *InsertLiveCuepointParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// InsertLiveCuepoint inserts a cuepoint into a live broadcast. The cuepoint
// triggers an ad break at the specified time in the broadcast. The broadcast
// must be actively streaming for the cuepoint to be inserted.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/liveCuepoints/insert
func InsertLiveCuepoint(runner RequestRunner, p *InsertLiveCuepointParams) (*LiveCuepoint, error) {
	body, err := jsonBody(p.Cuepoint)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    LiveCuepointsUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out LiveCuepoint
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
