package dto

import (
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
	"net/url"
)

type AuthCodeReq struct {
	UID         int    `json:"uid"`
	APPID       string `json:"appid"`
	Scope       string `json:"scope"`        //预留参数，后面需要的时间给
	RedirectUri string `json:"redirect_uri"` //要跳转的域名
}

func (a AuthCodeReq) Check() error {
	if a.UID <= 0 || len(a.APPID) != 10 || a.Scope == "" {
		return hcode.ParameterErr
	}
	return nil
}

func (a AuthCodeReq) GetRedirectUriHost() (string, error) {
	URL, err := url.Parse(a.RedirectUri)
	if err != nil {
		return "", hcode.ParameterErr
	}
	return URL.Host, nil
}
