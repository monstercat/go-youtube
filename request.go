package youtube

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	BaseUrlV3        = "https://www.googleapis.com/youtube/v3"
	YoutubePartnerV1 = "https://www.googleapis.com/youtube/partner/v1"
)

type RequestRunner interface {
	Run(r *Request) (*http.Response, error)
}

type Request struct {
	Method string
	Url    string
	Params url.Values
	Body   io.Reader
}

func DecodeResponse(res *http.Response, out interface{}) error {
	if res.StatusCode >= 400 {
		var e Error
		if err := json.NewDecoder(res.Body).Decode(&e); err == nil {
			e.StatusCode = res.StatusCode
			return e
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		e.StatusCode = res.StatusCode
		e.ErrorType = ErrTypeUnknown
		e.Description = string(body)
		return e
	}

	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		body, bodyErr := ioutil.ReadAll(res.Body)
		if bodyErr != nil {
			return Error{
				ErrorType:   ErrTypeBody,
				Description: bodyErr.Error(),
			}
		}
		return Error{
			ErrorType:   ErrTypeJSON,
			Description: err.Error(),
			Body:        string(body),
		}
	}

	return nil
}