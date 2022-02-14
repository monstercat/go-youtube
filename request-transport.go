package youtube

import (
	"net/http"
)

// CustomClientRunner takes in a specialized http.Client and uses it to run the requests. This is useful, for example,
// for using a JWT token as an authentication parameter instead of an access token or an api key.
//
// e.g.,
//    conf, _ := google.JWTConfigFromJSON(jsonBytes, scopes...)
//    client := conf.Client(context.Background())
//
//    runner := &CustomClientRunner{ Client: client }
type CustomClientRunner struct {
	Client *http.Client
}

func (runner *CustomClientRunner) Run(r *Request) (*http.Response, error) {
	// TODO: do we need this?
	//values := url.Values{}
	//values.Set("alt", "json")
	//values.Set("prettyPrint", "false")
	req, err := http.NewRequest(r.Method, r.Url+"?"+r.Params.Encode(), r.Body)
	if err != nil {
		return nil, err
	}
	return runner.Client.Do(req)
}