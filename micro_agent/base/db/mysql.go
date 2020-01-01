package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"micro_agent/base/config"
)

func initMysql()  {
	if mysqlDb, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8",
		config.GetMysqlConfig().GetName(), config.GetMysqlConfig().GetPass(), config.GetMysqlConfig().GetIp(), config.GetMysqlConfig().GetIp(), config.GetMysqlConfig().GetPort()));err != nil {
		panic(err)
	}
	mysqlDb.SetMaxOpenConns(config.GetMysqlConfig().GetMaxIdle())
	mysqlDb.SetMaxIdleConns(config.GetMysqlConfig().GetMaxOpen())
	_ = mysqlDb.Ping()
	if config.GetServerConfig().AppIsDebug(){
		mysqlDb.ShowSQL(true)
		mysqlDb.ShowExecTime(true)
	}
	//tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "im_")
	//mysqlDb.SetTableMapper(tbMapper)
}

func closeMysql()  {
	if mysqlDb!=nil{
		_ = mysqlDb.Close()
	}
}