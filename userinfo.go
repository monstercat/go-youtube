package youtube

import (
	"net/http"
	"time"
)

const (
	UserInfoUrl = "https://www.googleapis.com/oauth2/v2/userinfo"
)

type UserInfo struct {
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
	Id            string `json:"id"`
	Hd            string `json:"hd"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
}

func GetUserInfo(token string, timeout time.Duration) (*UserInfo, error) {
	runner := &AccessTokenRunner{
		Timeout:     timeout,
		AccessToken: token,
	}
	res, err := runner.Run(&Request{
		Method:      http.MethodGet,
		Url:         UserInfoUrl,
	})
	if err != nil {
		return nil, err
	}

	var u UserInfo
	if err := DecodeResponse(res, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
