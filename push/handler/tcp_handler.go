package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var pushChan = make(chan bool, 50)



func ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method!="POST"{
		w.Write([]byte("请使用POST方法"))
		return
	}
	pushChan<-true
	defer func() {
		<-pushChan
	}()
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	var report ServerReport
	if err := json.Unmarshal(reqData, &report); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if report.Uid==0||report.Msg==""||report.Sign==""{
		w.Write([]byte("参数不完整"))
		return
	}
	if err:=PushMsg(report.Uid,report.Msg);err!=nil{
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("ok"))
	return

}
