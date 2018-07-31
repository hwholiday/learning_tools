package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/cihub/seelog"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/go-xorm/core"
)

type dbInfo struct {
	User string `json:"mysqluser"`
	Pass string `json:"mysqlpass"`
	Url  string `json:"mysqlurl"`
	Port  string `json:"mysqlport"`
	DbName   string `json:"mysqldb"`
}
var db *xorm.Engine

func InitDb(path string) {
	var err error
	info:=getDbInfoByPath(path)
	checkErr(err)
	sql:=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8",info.User, info.Pass,info.Url,info.Port,info.DbName)
	fmt.Println(sql)
	db, err = xorm.NewEngine("mysql", sql)
	checkErr(err)
	checkErr(db.Ping())
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "control_")
	db.SetTableMapper(tbMapper)
	db.ShowSQL(true)
}
func getDbInfoByPath(path string)(db dbInfo){
	data,err:=ioutil.ReadFile(path)
	checkErr(err)
	err=json.Unmarshal(data,&db)
	checkErr(err)
	return
}

func checkErr(err error){
	if err != nil {
		seelog.Error(err)
		seelog.Flush()
		os.Exit(0)
	}
}
