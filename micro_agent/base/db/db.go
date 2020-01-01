package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"sync"
)

var (
	err error
	mysqlDb *xorm.Engine
	m sync.Mutex
)

func Init()  {
	m.Lock()
	defer m.Unlock()
	initMysql()
}

func GetMySqlDb()*xorm.Engine {
	return mysqlDb
}

func CloseMySqlDb(){
	closeMysql()
}