package youtube

import (
	"net/http"
	"time"
)

type AccessTokenRunner struct {
	AccessToken string
	Timeout     time.Duration
}

func (runner *AccessTokenRunner) Run(r *Request) (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.Url+"?"+r.Params.Encode(), r.Body)
	if err != nil {
		return nil, err
	}
	if runner.AccessToken != "" {
		req.Header.Add("Authorization", "Bearer "+runner.AccessToken)
	}

	client := http.Client{
		Timeout: runner.Timeout,
	}
	return client.Do(req)
}