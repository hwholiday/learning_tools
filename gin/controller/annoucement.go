package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"test/gin/model"
)
//announcement
type AnnouncementController struct {
	BaseController
}

//添加公告
func (c *AnnouncementController) Add(g *gin.Context) {
	var data model.Announcement
	err := g.BindJSON(&data)
	if err != nil {
		c.ResponseFailureForParameter(g, CErrJSON)
		return
	}
	if err := model.AddAnnouncement(&data); err != nil {
		c.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	c.ResponseSuccess(g)
	return
}

//删除公告
func (c *AnnouncementController) Del(g *gin.Context) {
	sId := g.Param("id")
	if sId == "" {
		c.ResponseFailureForParameter(g, CErrParam)
		return
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		c.ResponseFailureForFuncErr(g, CErrTypeConversion)
	}
	if err := model.DeleteAnnouncement(&model.Announcement{Id: id}); err != nil {
		c.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	c.ResponseSuccess(g)
	return
}

//修改公告
func (c *AnnouncementController) UpDate(g *gin.Context) {
	sId := g.Param("id")
	if sId == "" {
		c.ResponseFailureForParameter(g, CErrParam)
		return
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		c.ResponseFailureForFuncErr(g, CErrTypeConversion)
	}
	var data model.Announcement
	err = g.BindJSON(&data)
	if err != nil {
		c.ResponseFailureForParameter(g, CErrJSON)
		return
	}
	data.Id = id
	data.UpdateTime = time.Now().Unix()
	if err := model.UpdateAnnouncement(&data); err != nil {
		c.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	c.ResponseSuccess(g)
	return
}

//获取公告
// @Summary 获取公告
// @Tags  announcement
// @Description 通过ID获取公告信息
// @Accept  json
// @Produce  json
// @Param id path int false "Announcement ID"
// @Success 200 {string} model.Announcement{}
// @Failure 500 {string} string
// @Router /v1/announcement/{id} [get]
func (c *AnnouncementController) GetById(g *gin.Context) {
	sId := g.Param("id")
	if sId == "" {
		c.ResponseFailureForParameter(g, CErrParam)
		return
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		c.ResponseFailureForFuncErr(g, CErrTypeConversion)
	}
	data,err := model.GetAnnouncementById(id)
	if err != nil {
		c.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	c.ResponseData(g,data)
	return
}

// @Summary 删除公告
// @Tags  announcement
// @Produce  json
// @Param id param int true "ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router /v1/announcement/{id} [delete]
func (c *AnnouncementController) GetAll(g *gin.Context) {
	data,err := model.GetAnnouncementAll()
	if err != nil {
		c.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	c.ResponseData(g,data)
	return
}
