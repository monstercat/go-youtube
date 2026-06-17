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
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// Content ID API error envelope:
		// {"error":{"code":..,"message":..,"errors":[{reason,domain,message}]}}.
		// Tried first because its object-valued "error" field cannot decode
		// into the OAuth shape below.
		var env struct {
			Error struct {
				Code    int           `json:"code"`
				Message string        `json:"message"`
				Errors  []ErrorDetail `json:"errors"`
			} `json:"error"`
		}
		if err := json.Unmarshal(body, &env); err == nil &&
			(env.Error.Code != 0 || len(env.Error.Errors) > 0) {
			e := Error{
				StatusCode:  res.StatusCode,
				ErrorType:   ErrTypeUnknown,
				Description: env.Error.Message,
				Body:        string(body),
				Errors:      env.Error.Errors,
			}
			if len(env.Error.Errors) > 0 {
				e.Reason = env.Error.Errors[0].Reason
			}
			return e
		}

		// OAuth-style error: {"error":"invalid_grant","error_description":".."}.
		var oauth struct {
			Error       ErrorType `json:"error"`
			Description string    `json:"error_description"`
		}
		if err := json.Unmarshal(body, &oauth); err == nil {
			return Error{
				StatusCode:  res.StatusCode,
				ErrorType:   oauth.Error,
				Description: oauth.Description,
				Body:        string(body),
			}
		}

		// Unrecognised shape (e.g. an HTML or plain-text body).
		return Error{
			StatusCode:  res.StatusCode,
			ErrorType:   ErrTypeUnknown,
			Description: string(body),
			Body:        string(body),
		}
	}
	if out == nil {
		return nil
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