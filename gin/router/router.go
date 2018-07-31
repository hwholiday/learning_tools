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
	announcement := &controller.AnnouncementController{}
	v1:=r.Group("/v1")
	{
		//公告相关业务
		v1.GET("/announcement/:id", announcement.GetById)
		v1.GET("/announcement/", announcement.GetAll)
		v1.POST("/announcement", announcement.Add)
		v1.PUT("/announcement/:id", announcement.UpDate)
		v1.DELETE("/announcement/:id", announcement.Del)
	}
}
