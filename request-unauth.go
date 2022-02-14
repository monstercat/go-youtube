package youtube

import (
	"net/http"
	"time"
)

type UnauthenticatedRunner struct {
	/// Timeout for the request
	Timeout time.Duration
}

func (u *UnauthenticatedRunner) Run(r *Request) (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.Url+"?"+r.Params.Encode(), r.Body)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: u.Timeout,
	}
	return client.Do(req)
}
