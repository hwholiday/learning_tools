package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetCorsRouters(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "lang", "token", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}
