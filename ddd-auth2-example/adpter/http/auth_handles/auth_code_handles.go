package auth_handles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/dto"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"go.uber.org/zap"
	"net/http"
)

// CreateCodeOpenId 参数	是否必须	说明
//appid	是	公众号的唯一标识
//redirect_uri	是	授权后重定向的回调链接地址， 请使用 urlEncode 对链接进行处理
//response_type	是	返回类型，请填写code
//scope	是	应用授权作用域，snsapi_base （不弹出授权页面，直接跳转，只能获取用户openid），snsapi_userinfo （弹出授权页面，可通过openid拿到昵称 ）
//state	否	重定向后会带上state参数，开发者可以填写a-zA-Z0-9的参数值，最多128字节

func (h *Handles) CreateCodeOpenId(g *gin.Context) {
	var (
		appid        = g.Query("appid")
		redirectUri  = g.Query("redirect_uri")
		responseType = g.Query("response_type")
		scope        = g.Query("scope")
		state        = g.Query("state")
	)
	if redirectUri == "" || scope == "" || state == "" || responseType == "" {
		h.ResponseErr(g, hcode.ParameterErr)
		return
	}
	//TODO  uid  这里的uid可以设计为用token置换来取
	var uid = 1
	code, err := h.auth.CreateCodeOpenId(g, dto.AuthCodeReq{
		UID:         uid,
		APPID:       appid,
		Scope:       scope,
		RedirectUri: redirectUri,
	})
	if err != nil {
		h.ResponseErr(g, err)
		return
	}
	//?code=CODE&state=STATE
	url := fmt.Sprintf("%s?code=%s&state=%s", redirectUri, code, state)
	log.GetLogger().Debug("CreateCodeOpenId", zap.Any("url", url))
	g.Redirect(http.StatusMovedPermanently, url)
	return

}
