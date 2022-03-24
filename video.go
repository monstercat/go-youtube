package youtube

import (
	"net/http"
	"net/url"
	"strings"
)

const (
	ListVideosUrl = BaseUrlV3 + "/videos"
)

// VideoSnippet - the snippet object contains basic details about the video,
// such as its title, description, and category.
type VideoSnippet struct {
	// CategoryId: The YouTube video category associated with the video.
	CategoryId string `json:"categoryId,omitempty"`

	// ChannelId: The ID that YouTube uses to uniquely identify the channel
	// that the video was uploaded to.
	ChannelId string `json:"channelId,omitempty"`

	// ChannelTitle: Channel title for the channel that the video belongs
	// to.
	ChannelTitle string `json:"channelTitle,omitempty"`

	// Description: The video's description. @mutable youtube.videos.insert
	// youtube.videos.update
	Description string `json:"description,omitempty"`

	// PublishedAt: The date and time when the video was uploaded.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Tags: A list of keyword tags associated with the video. Tags may
	// contain spaces.
	Tags []string `json:"tags,omitempty"`

	// Title: The video's title. @mutable youtube.videos.insert
	// youtube.videos.update
	Title string `json:"title,omitempty"`
}

type Video struct {
	// Id: The ID that YouTube uses to uniquely identify the video.
	Id string `json:"id,omitempty"`

	// TODO: fill in the other parts as well.

	// Snippet: The snippet object contains basic details about the video,
	// such as its title, description, and category.
	Snippet *VideoSnippet
}

type ListVideosResponse struct {
	// Items are the returned videos
	Items []*Video `json:"items"`

	// NextPageToken: The token that can be used as the value of the
	// pageToken parameter to retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken"`

	// PageInfo: General pagination information.
	PageInfo *PageInfo `json:"pageInfo"`
}

type ListVideoParamsPart string

const (
	ListVideoParamsPartContentDetails       ListVideoParamsPart = "contentDetails"
	ListVideoParamsPartFileDetails          ListVideoParamsPart = "fileDetails"
	ListVideoParamsPartId                   ListVideoParamsPart = "id"
	ListVideoParamsPartLiveStreamingDetails ListVideoParamsPart = "liveStreamingDetails"
	ListVideoParamsPartLocalizations        ListVideoParamsPart = "localizations"
	ListVideoParamsPartPlayer               ListVideoParamsPart = "player"
	ListVideoParamsPartProcessingDetails    ListVideoParamsPart = "processingDetails"
	ListVideoParamsPartRecordingDetails     ListVideoParamsPart = "recordingDetails"
	ListVideoParamsPartSnippet              ListVideoParamsPart = "snippet"
	ListVideoParamsPartStatistics           ListVideoParamsPart = "statistics"
	ListVideoParamsPartStatus               ListVideoParamsPart = "status"
	ListVideoParamsPartSuggestions          ListVideoParamsPart = "suggestions"
	ListVideoParamsPartTopicDetails         ListVideoParamsPart = "topicDetails"
)

func (p ListVideoParamsPart) Valid() bool {
	switch p {
	case ListVideoParamsPartContentDetails,
		ListVideoParamsPartFileDetails,
		ListVideoParamsPartId,
		ListVideoParamsPartLiveStreamingDetails,
		ListVideoParamsPartLocalizations,
		ListVideoParamsPartPlayer,
		ListVideoParamsPartProcessingDetails,
		ListVideoParamsPartRecordingDetails,
		ListVideoParamsPartSnippet,
		ListVideoParamsPartStatistics,
		ListVideoParamsPartStatus,
		ListVideoParamsPartSuggestions,
		ListVideoParamsPartTopicDetails:
		return true
	}
	return false
}

type ListVideoParams struct {
	// Part parameter specifies a comma-separated list of one or more video resource properties that the API
	// response will include. This is required.
	Parts []ListVideoParamsPart

	// Ids to filter by. Optional
	Ids []string

	// The PageToken parameter identifies a specific page in the result set that should be returned. In an API
	// response, the nextPageToken and prevPageToken properties identify other pages that could be retrieved.
	//
	// Note: This parameter is supported for use in conjunction with the myRating parameter, but it is not supported
	// for use in conjunction with the id parameter.
	PageToken string
}

func (o *ListVideoParams) convertParts() []string {
	if len(o.Parts) == 0 {
		return []string{}
	}
	p := make([]string, 0, len(o.Parts))
	for _, prompt := range o.Parts {
		p = append(p, string(prompt))
	}
	return p
}

func (o *ListVideoParams) Values() url.Values {
	vals := url.Values{}
	vals.Add("part", strings.Join(o.convertParts(), ","))
	if o.PageToken != "" {
		vals.Add("pageToken", o.PageToken)
	} else if len(o.Ids) > 0 {
		vals.Add("id", strings.Join(o.Ids, ","))
	}
	return vals
}

// ListVideos provides a list of videos to the user.
//
// This function has a quota cost of 1 unit.
// https://developers.google.com/youtube/v3/docs/videos/list
func ListVideos(runner RequestRunner, p *ListVideoParams) (*ListVideosResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    ListVideosUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out ListVideosResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
