package base

import (
	"micro_agent/base/config"
	"micro_agent/base/db"
	"micro_agent/base/tool"
)

func Init(path string) {
	config.Init(path)
	db.Init()
	tool.Init()
}
