package dto

import "github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"

type AuthTokenSimple struct {
	OpenID       string `json:"open_id"`
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int64  `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

type ProduceAuthTokenReq struct {
	Code      string `json:"code"`
	APPID     string `json:"appid"`
	Secret    string `json:"secret"`
	GrantType string `json:"grant_type"` //填写为authorization_code
}

type RefreshAccessTokenReq struct {
	APPID        string `json:"appid"`
	GrantType    string `json:"grant_type"`    //填写为 refresh_token
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
}

func (a RefreshAccessTokenReq) Check() error {
	if a.APPID == "" || a.GrantType == "" || a.RefreshToken == "" {
		return hcode.ParameterErr
	}
	return nil
}

type OpenIdTokenReq struct {
	OpenId string `json:"open_id"`
	Token  string `json:"token"`
}

func (a OpenIdTokenReq) Check() error {
	if a.OpenId == "" || a.Token == "" {
		return hcode.ParameterErr
	}
	return nil
}
