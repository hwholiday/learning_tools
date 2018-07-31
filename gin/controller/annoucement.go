package controller

import (
	"github.com/gin-gonic/gin"
	"bat_server/bat_messager/bat_gw/model"
)

type AnnouncementController struct {
	BaseController
}

func (c *AnnouncementController) Add(g *gin.Context) {
	var data model.Announcement
	err := g.BindJSON(&data)
	if err != nil {
		c.ResponseFailureForParameter(g, "获取json数据失败")
		return
	}
	if err := model.AddAnnouncement(&data); err != nil {
		c.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	c.ResponseSuccess(g)
	return
}
