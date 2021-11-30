package obj

import (
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/dto"
	"time"
)

type AuthToken struct {
	APPID                string `json:"appid"`
	Secret               string `json:"secret"`
	OpenID               string `json:"open_id"`
	AccessToken          string `json:"access_token"`           // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	AccessTokenTimeline  int64  `json:"access_token_timeline"`  // access_token 到期时间毫秒
	RefreshToken         string `json:"refresh_token"`          // 用户刷新access_token
	RefreshTokenTimeline int64  `json:"refresh_token_timeline"` // refresh_token 到期时间毫秒
	Scope                string `json:"scope"`                  // 用户授权的作用域，使用逗号（,）分隔
}

func (a AuthToken) TOSimple() dto.AuthTokenSimple {
	return dto.AuthTokenSimple{
		OpenID:       a.OpenID,
		AccessToken:  a.AccessToken,
		ExpiresIn:    a.GetExpiresIn(),
		RefreshToken: a.RefreshToken,
		Scope:        a.Scope,
	}
}

func (a AuthToken) GetExpiresIn() int64 {
	now := time.Now().Unix()
	t := a.AccessTokenTimeline - now
	if t > 0 {
		return t
	}
	return 0
}

func (a AuthToken) AccessTokenExpired() bool {
	now := time.Now().Unix()
	t := a.AccessTokenTimeline - now
	if t > 0 {
		return false
	}
	return true
}

func (a AuthToken) RefreshTokenExpired() bool {
	now := time.Now().Unix()
	t := a.RefreshTokenTimeline - now
	if t > 0 {
		return false
	}
	return true
}
