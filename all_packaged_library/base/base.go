package base

import (
	"file_storage/base/config"
	"file_storage/base/db"
	"file_storage/base/tool"
)
//配置文件的目录
func Init(path string) {
	config.Init(path)
	tool.Init()
	db.Init()
}

