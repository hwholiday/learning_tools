package routers

import (
	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/adpter/http/auth_handles"
)

func SetRouters(r *gin.Engine, h *auth_handles.Handles) {
	SetCorsRouters(r)
	r.GET("/oauth2/authorize", h.CreateCodeOpenId)       //获取code
	r.GET("/oauth2/access_token", h.CreateToken)         //通过code换取网页授权access_token
	r.GET("/oauth2/refresh_token", h.RefreshAccessToken) //刷新access_token
	r.GET("/oauth2/userinfo", h.UserInfo)                //拉取用户信息(需scope为 snsapi_userinfo)
	r.GET("/oauth2/auth", h.CheckAccessToken)            //检验授权凭证（access_token）是否有效
}
