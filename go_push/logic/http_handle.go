package logic

import (
	"fmt"
	"github.com/hwholiday/learning_tools/go_push/gateway"
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
	gateway.GetPushManage().Push(&gateway.PushJob{
		Type:     1,
		PushType: pushType,
		Info:     val,
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
	gateway.GetPushManage().Push(&gateway.PushJob{
		Type:     2,
		RoomId:   roomId,
		PushType: pushType,
		Info:     val,
	})
	_, _ = res.Write([]byte(fmt.Sprintf("[%s]房间推送任务添加成功", gateway.RoomTitle[roomId])))
	return
}

func HttpRoomJoin(res http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	wsId := req.FormValue("wsId")
	id := req.FormValue("id")
	if wsId == "" || id == "" {
		_, _ = res.Write([]byte("值为空"))
		return
	}
	roomId, _ := strconv.Atoi(id)
	err := gateway.GetRoomManage().AddRoom(roomId, wsId)
	if err != nil {
		_, _ = res.Write([]byte(err.Error()))
		return
	}
	_, _ = res.Write([]byte(fmt.Sprintf("加入[%s]房间成功", gateway.RoomTitle[roomId])))
	return
}

func HttpRoomLeave(res http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	wsId := req.FormValue("wsId")
	id := req.FormValue("id")
	if wsId == "" || id == "" {
		_, _ = res.Write([]byte("值为空"))
		return
	}
	roomId, _ := strconv.Atoi(id)
	err := gateway.GetRoomManage().LeaveRoom(roomId, wsId)
	if err != nil {
		_, _ = res.Write([]byte(err.Error()))
		return
	}
	_, _ = res.Write([]byte(fmt.Sprintf("离开[%s]房间成功", gateway.RoomTitle[roomId])))
	return
}
