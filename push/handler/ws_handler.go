package handler

import (
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
	"log"
)

type connection struct {
	ws   *websocket.Conn
	send chan []byte
	uid  int
}

//接受消息
func (c *connection) reader() {
	for {
		messageType, msg, err := c.ws.ReadMessage()
		if err != nil {
			log.Println("ws_handler_23:"+err.Error())
			log.Println("ws_handler_24:"+strconv.Itoa(messageType))
			return
		}
		var clients ClientsReport
		if err := json.Unmarshal(msg, &clients); err != nil {
			log.Println("ws_handler_28:"+err.Error())
			c.ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			return
		}
		if clients.Status==1{
			c.ws.WriteMessage(websocket.TextMessage, []byte("ok"))
		}else {
			H.broadcast <- &clients
		}
	}
	c.ws.Close()
}

//发出消息

func (c *connection) writer() {
	for msg := range c.send {
		if err := c.ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("ws_handler_46:"+err.Error())
			return
		}
	}
	c.ws.Close()

}

var upgrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024}

func PushHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	strUid := req.Form.Get("uid")
	strTime := req.Form.Get("time")
	rid := req.Form.Get("sign")
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("ws_handler_68:"+err.Error())
		respWs(ws,err.Error())
		return
	}
	if len(strUid) == 0 || len(rid) == 0 || len(strTime) == 0 {
		respWs(ws,"参数不完整")
		return
	}
	id, err := strconv.Atoi(strUid)
	if err!=nil{
		log.Println("ws_handler_78:"+err.Error())
		respWs(ws,err.Error())
		return
	}
	upTime, err := strconv.ParseInt(strTime, 10, 64)
	if err!=nil{
		log.Println(err.Error())
		respWs(ws,err.Error())
		return
	}
	endTime:=time.Now().Unix()+60
	startTime:=time.Now().Unix()-60
    if  upTime<startTime||upTime>endTime{
		respWs(ws,"校验码在获取后60秒内有效")
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws, uid: id}
	H.register <- c
	defer func() { H.unregister <- c }()
	go c.writer()
	c.reader()
}

func respWs(ws *websocket.Conn,data string)  {
	ws.WriteMessage(websocket.TextMessage,[]byte(data))
	ws.Close()
}
