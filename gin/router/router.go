package router

import (
	"github.com/gin-gonic/gin"
	"bat_server/bat_messager/bat_gw/controller"
	"github.com/gin-contrib/cors"
)

func SetRouters(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	user := &controller.UserController{}
	r.GET("/user/version", user.Version)
}
