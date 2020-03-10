package gateway

import (
	"net/http"
	"strconv"
)

func HttpPushAll(res http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	val := req.FormValue("val")
	tag := req.FormValue("tag")
	if val == "" || tag == "" {
		_, _ = res.Write([]byte("值为空"))
		return
	}
	pushType, _ := strconv.Atoi(tag)
	GetPushManage().Push(&PushJob{
		Type:     1,
		PushType: pushType,
		info:     val,
	})
	_, _ = res.Write([]byte("全部推送任务添加成功"))
	return
}

func HttpPushRoom(res http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	val := req.FormValue("val")
	tag := req.FormValue("tag")
	id := req.FormValue("id")
	if val == "" || tag == "" || id == "" {
		_, _ = res.Write([]byte("值为空"))
		return
	}
	pushType, _ := strconv.Atoi(tag)
	roomId, _ := strconv.Atoi(id)
	GetPushManage().Push(&PushJob{
		Type:     2,
		roomId:   roomId,
		PushType: pushType,
		info:     val,
	})
	_, _ = res.Write([]byte("房间推送任务添加成功"))
	return
}
