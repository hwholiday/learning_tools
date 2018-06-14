package handler

//服务端上报数据
type ServerReport struct {
	Uid  int    `json:"uid"`
	Msg  string `json:"msg"`
	Sign string `json:"sign"`
}

//客户端上报
type ClientsReport struct {
	Uid    int `json:"uid"`
	Status int `json:"status"` //1心跳包,2其他
	Msg  string `json:"msg"`
}
