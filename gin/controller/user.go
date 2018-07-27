package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/cihub/seelog"
)

type UserController struct{}

func (u *UserController) Version(g *gin.Context) {
	seelog.Info("success")
	g.String(200, "ok")
}
