package youtube

import (
	"net/http"
	"net/url"
)

const (
	MetadataHistoryUrl = YoutubePartnerV1 + "/metadataHistory"
)

// Metadata contains comprehensive metadata for an asset, including
// identifiers (ISRC, ISWC, UPC), descriptive fields (title, artist,
// genre), and categorization data.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assets#Metadata
type Metadata struct {
	// Actor is a list of actors or performers in the content.
	Actor []string `json:"actor,omitempty"`
	// Album is the album name for a music asset.
	Album string `json:"album,omitempty"`
	// Artist is a list of artists or performers for the content.
	Artist []string `json:"artist,omitempty"`
	// Broadcaster is a list of broadcasters for the content.
	Broadcaster []string `json:"broadcaster,omitempty"`
	// Category is the content category, such as "entertainment" or "news".
	Category string `json:"category,omitempty"`
	// ContentType is the type of content (e.g., "movie", "episode", "song").
	ContentType string `json:"contentType,omitempty"`
	// CopyrightDate is the copyright date of the content.
	CopyrightDate *Date `json:"copyrightDate,omitempty"`
	// CustomId is a unique value that you, the metadata provider, use to
	// identify an asset. The value can be up to 64 bytes.
	CustomId string `json:"customId,omitempty"`
	// Description is a description of the content.
	Description string `json:"description,omitempty"`
	// Director is a list of directors for the content.
	Director []string `json:"director,omitempty"`
	// Eidr is the Entertainment Identifier Registry (EIDR) assigned to the asset.
	Eidr string `json:"eidr,omitempty"`
	// EndYear is the last year in which the content was or will be released.
	EndYear int `json:"endYear,omitempty"`
	// EpisodeNumber is the episode number for an episodic asset.
	EpisodeNumber string `json:"episodeNumber,omitempty"`
	// EpisodesAreUntitled indicates whether the episodes in the show are untitled.
	EpisodesAreUntitled bool `json:"episodesAreUntitled,omitempty"`
	// Genre is a list of genres that describe the content.
	Genre []string `json:"genre,omitempty"`
	// Grid is the Global Release Identifier (GRid) of a music asset.
	Grid string `json:"grid,omitempty"`
	// Hfa is the Harry Fox Agency (HFA) song code for a music asset.
	Hfa string `json:"hfa,omitempty"`
	// InfoUrl is a URL providing more information about the content.
	InfoUrl string `json:"infoUrl,omitempty"`
	// Isan is the International Standard Audiovisual Number (ISAN) for the content.
	Isan string `json:"isan,omitempty"`
	// Isrc is the International Standard Recording Code (ISRC) of a music
	// or sound recording asset.
	Isrc string `json:"isrc,omitempty"`
	// Iswc is the International Standard Musical Work Code (ISWC) of a
	// musical composition asset.
	Iswc string `json:"iswc,omitempty"`
	// Keyword is a list of keywords associated with the content.
	Keyword []string `json:"keyword,omitempty"`
	// Label is the record label that released a sound recording asset.
	Label string `json:"label,omitempty"`
	// Notes is free-form notes about the asset.
	Notes string `json:"notes,omitempty"`
	// OriginalReleaseMedium is the original release medium (e.g., "Film",
	// "Digital", "CD").
	OriginalReleaseMedium string `json:"originalReleaseMedium,omitempty"`
	// Producer is a list of producers for the content.
	Producer []string `json:"producer,omitempty"`
	// Ratings is a list of content ratings for the asset, such as "PG" or "R".
	Ratings []*Rating `json:"ratings,omitempty"`
	// ReleaseDate is the date that the content was first publicly released.
	ReleaseDate *Date `json:"releaseDate,omitempty"`
	// SeasonNumber is the season number for an episodic asset.
	SeasonNumber string `json:"seasonNumber,omitempty"`
	// ShowCustomId is a custom ID that you use to identify the show to which
	// this episode belongs.
	ShowCustomId string `json:"showCustomId,omitempty"`
	// ShowTitle is the title of the show to which an episode belongs.
	ShowTitle string `json:"showTitle,omitempty"`
	// SpokenLanguage is the spoken language of the content.
	SpokenLanguage string `json:"spokenLanguage,omitempty"`
	// StartYear is the first year in which the content was or will be released.
	StartYear int `json:"startYear,omitempty"`
	// SubtitledLanguage is a list of languages in which subtitles are available.
	SubtitledLanguage []string `json:"subtitledLanguage,omitempty"`
	// Title is the title of the asset.
	Title string `json:"title,omitempty"`
	// TmsId is the Tribune Media Systems (TMS) ID of the asset.
	TmsId string `json:"tmsId,omitempty"`
	// TotalEpisodesExpected is the total number of full-length episodes in
	// the season. Applicable only to episodic assets.
	TotalEpisodesExpected int `json:"totalEpisodesExpected,omitempty"`
	// Upc is the Universal Product Code (UPC) of the asset.
	Upc string `json:"upc,omitempty"`
	// Writer is a list of writers or songwriters for the content.
	Writer []string `json:"writer,omitempty"`
}

// MetadataHistory provides a snapshot of asset metadata at a specific point
// in time, along with information about who provided the metadata.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/metadataHistory#MetadataHistory
type MetadataHistory struct {
	// Kind identifies the API resource type. The value is "youtubePartner#metadataHistory".
	Kind string `json:"kind,omitempty"`
	// Metadata contains the metadata values at this point in the asset's history.
	Metadata *Metadata `json:"metadata,omitempty"`
	// Origination contains information about the metadata provider, including
	// the content owner and the source.
	Origination *Origination `json:"origination,omitempty"`
	// TimeProvided is the date and time the metadata was provided.
	TimeProvided string `json:"timeProvided,omitempty"`
}

// MetadataHistoryListResponse is the response for metadataHistory.list.
type MetadataHistoryListResponse struct {
	// Items is a list of metadata history entries for the asset.
	Items []*MetadataHistory `json:"items"`
	// Kind identifies the API resource type. The value is "youtubePartner#metadataHistoryList".
	Kind string `json:"kind,omitempty"`
}

// ── metadataHistory.list ────────────────────────────────────────────────────

// ListMetadataHistoryParams are parameters for metadataHistory.list.
type ListMetadataHistoryParams struct {
	// AssetId is the ID of the asset for which to retrieve metadata history.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *ListMetadataHistoryParams) Values() url.Values {
	v := url.Values{}
	if p.AssetId != "" {
		v.Set("assetId", p.AssetId)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// ListMetadataHistory retrieves the metadata history for the specified asset.
// Each entry in the result set represents a snapshot of the asset's metadata
// at a specific point in time, along with information about who provided it.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/metadataHistory/list
func ListMetadataHistory(runner RequestRunner, p *ListMetadataHistoryParams) (*MetadataHistoryListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    MetadataHistoryUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out MetadataHistoryListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
