package auth_handles

import (
	"github.com/gin-gonic/gin"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/dto"
)

// CreateToken appid	是	唯一标识
//secret	是	secret
//code	是	填写第一步获取的code参数
//grant_type	是	填写为authorization_code
func (h *Handles) CreateToken(g *gin.Context) {
	var (
		appid     = g.Query("appid")
		secret    = g.Query("secret")
		code      = g.Query("code")
		grantType = g.Query("grant_type")
	)
	auth, err := h.auth.CreateToken(g, dto.ProduceAuthTokenReq{
		Code:      code,
		APPID:     appid,
		Secret:    secret,
		GrantType: grantType,
	})
	if err != nil {
		h.ResponseErr(g, err)
		return
	}
	h.ResponseData(g, auth)
	return
}

// RefreshAccessToken  appid	是	唯一标识
//grant_type	是	填写为 refresh_token
//refresh_token	是	填写通过access_token获取到的refresh_token参数
func (h *Handles) RefreshAccessToken(g *gin.Context) {
	var (
		appid        = g.Query("appid")
		grantType    = g.Query("grant_type")
		refreshToken = g.Query("refresh_token")
	)
	auth, err := h.auth.RefreshAccessToken(g, dto.RefreshAccessTokenReq{
		APPID:        appid,
		GrantType:    grantType,
		RefreshToken: refreshToken,
	})
	if err != nil {
		h.ResponseErr(g, err)
		return
	}
	h.ResponseData(g, auth)
	return
}

func (h *Handles) CheckAccessToken(g *gin.Context) {
	var (
		openId      = g.Query("openid")
		accessToken = g.Query("access_Token")
	)
	err := h.auth.CheckToken(g, dto.OpenIdTokenReq{
		OpenId: openId,
		Token:  accessToken,
	})
	if err != nil {
		h.ResponseErr(g, err)
		return
	}
	h.ResponseSuccess(g)
	return
}

func (h *Handles) UserInfo(g *gin.Context) {
	var (
		openId      = g.Query("openid")
		accessToken = g.Query("access_Token")
	)
	data, err := h.auth.GetUserInfo(g, dto.OpenIdTokenReq{
		OpenId: openId,
		Token:  accessToken,
	})
	if err != nil {
		h.ResponseErr(g, err)
		return
	}
	h.ResponseData(g, data)
	return
}
