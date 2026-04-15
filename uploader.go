package youtube

import (
	"net/http"
	"net/url"
)

const (
	UploaderUrl = YoutubePartnerV1 + "/uploader"
)

// Uploader represents an uploader account that can be used to upload content
// delivery packages.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/uploader#Uploader
type Uploader struct {
	// Kind identifies the API resource type. The value is "youtubePartner#uploader".
	Kind string `json:"kind,omitempty"`
	// UploaderName is the name of the uploader account. This name is used
	// when uploading content delivery packages.
	UploaderName string `json:"uploaderName,omitempty"`
}

// UploaderListResponse is the response for uploader.list.
type UploaderListResponse struct {
	// Items is a list of uploader accounts that match the request criteria.
	Items []*Uploader `json:"items"`
	// Kind identifies the API resource type. The value is "youtubePartner#uploaderList".
	Kind string `json:"kind,omitempty"`
}

// ── uploader.list ───────────────────────────────────────────────────────────

// ListUploadersParams are parameters for uploader.list.
type ListUploadersParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *ListUploadersParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// ListUploaders retrieves a list of uploader accounts for the content owner.
// Uploaders are accounts that can be used to upload content delivery packages
// via SFTP or the Aspera dropbox.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/uploader/list
func ListUploaders(runner RequestRunner, p *ListUploadersParams) (*UploaderListResponse, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    UploaderUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out UploaderListResponse
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
