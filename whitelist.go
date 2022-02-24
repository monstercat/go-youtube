package youtube

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrNotWhitelisted = errors.New("not whitelisted")
	ErrRateLimited    = errors.New("rate limited")
)

const (
	WhitelistUrl = YoutubePartnerV1 + "/whitelists"
)

// Whitelist represents a resource that can be whitelisted. This is technically a YouTube Channel (e.g., the
// Id will be equivalent to a YouTube Channel ID and the Title will be equivalent to a YouTube Channel Title.
// https://developers.google.com/youtube/partner/docs/v1/whitelists#resource
type Whitelist struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// GetWhitelistParams are parameters specifically for the GetWhitelist method.
type GetWhitelistParams struct {
	// Whitelist Id to retrieve
	Id string

	// The content owner that we are requesting whitelist on behalf of.
	OnBehalfOfContentOwner string
}

func (p *GetWhitelistParams) Values() url.Values {
	vals := url.Values{}
	vals.Add("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	return vals
}

// GetWhitelist returns a whitelist for a specific channel ID. It will return a resource if the channel is whitelisted
// and return an error if not whitelisted.
// https://developers.google.com/youtube/partner/docs/v1/whitelists/get
func GetWhitelist(runner RequestRunner, p *GetWhitelistParams) (*Whitelist, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    WhitelistUrl + "/" + p.Id,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}

	var out Whitelist
	if err := DecodeResponse(res, &out); err != nil {
		return nil, convertWhitelistError(err)
	}
	return &out, nil
}

type InsertWhitelistParams struct {
	// Channel Id to insert
	Id string

	// The content owner that we are requesting whitelist on behalf of.
	OnBehalfOfContentOwner string
}

func (p *InsertWhitelistParams) Body() (io.Reader, error) {
	m := make(map[string]interface{})
	m["kind"] = "youtubePartner#whitelist"
	m["id"] = p.Id

	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (p *InsertWhitelistParams) Values() url.Values {
	vals := url.Values{}
	vals.Add("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	return vals
}

// InsertWhitelist  - Whitelists a YouTube channel for your content owner. Whitelisted channels are channels that are
// not owned or managed by you, but you would like to whitelist so that no claims from your assets are placed on videos
// uploaded to these channels.
// https://developers.google.com/youtube/partner/docs/v1/whitelists/insert
func InsertWhitelist(runner RequestRunner, p *InsertWhitelistParams) (*Whitelist, error) {
	body, err := p.Body()
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    WhitelistUrl,
		Body:   body,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}

	var out Whitelist
	if err := DecodeResponse(res, &out); err != nil {
		return nil, convertWhitelistError(err)
	}
	return &out, nil
}

type DeleteWhitelistParams struct {
	// Channel Id to insert
	Id string

	// The content owner that we are requesting whitelist on behalf of.
	OnBehalfOfContentOwner string
}

func (p *DeleteWhitelistParams) Values() url.Values {
	vals := url.Values{}
	vals.Add("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	return vals
}

// DeleteWhitelist - Removes a whitelisted channel for a content owner
func DeleteWhitelist(runner RequestRunner, p *DeleteWhitelistParams) error {
	res, err := runner.Run(&Request{
		Method: http.MethodDelete,
		Url:    WhitelistUrl + "/" + p.Id,
		Params: p.Values(),
	})
	if err != nil {
		return err
	}
	return convertWhitelistError(DecodeResponse(res, nil))
}

// Converts whitelist errors according to https://developers.google.com/youtube/partner/docs/v1/errors#general
// NOTE: for the GetWhitelist route, was unable to get invalidValue & required errors.
// NOTE: for the InsertWhitelist route, was unable to get any errors including the channelNotFound error.
//
// quotaExceeded usually shows as a 403 error.
func convertWhitelistError(err error) error {
	if err == nil {
		return nil
	}
	v, ok := err.(*Error)
	if !ok {
		return err
	}
	if v.StatusCode == 404 {
		return ErrNotWhitelisted
	}
	if v.StatusCode == 403 {
		if strings.Index(v.Body, "quotaExceeded") > -1 {
			return ErrRateLimited
		}
	}
	return err
}
