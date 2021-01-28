package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"learning_tools/istio/api"
)

type Data struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func wsHandle(w *WsConnection) {
	var (
		err error
		msg []byte
	)
	fmt.Println("开启链接: ", w.GetIp(), "UID", w.GetUid(), "WSID", w.GetWsId())
	for {
		if msg, err = w.ReadMsg(); err != nil {
			fmt.Println("ReadMsg", err)
			wsClose(w)
			return
		}
		var data Data
		if err = json.Unmarshal(msg, &data); err != nil {
			fmt.Println("Unmarshal", err)
			continue
		}
		conn, err := grpc.Dial("logic.im.svc.cluster.local:8099", grpc.WithInsecure())
		if err != nil {
			fmt.Println("Dial", err)
			continue
		}
		ctx := context.WithValue(context.Background(), "data", data.Data)
		c := api.NewNameClient(conn)
		switch data.Code {
		case 1:
			info, err := c.ReqVersion(ctx, &api.Req{})
			if err != nil {
				fmt.Println("ReqVersion", err)
				fmt.Println(err)
				continue
			}
			d, _ := json.Marshal(&Data{
				Code: data.Code,
				Data: info.Name,
			})
			_ = w.SendMsg(d)
			fmt.Println("ReqVersion", info.Name)
		case 2:
			info, err := c.ReqName(ctx, &api.Req{
				Name: data.Data,
			})
			if err != nil {
				fmt.Println("ReqName", err)
				continue
			}
			d, _ := json.Marshal(&Data{
				Code: data.Code,
				Data: info.Name,
			})
			_ = w.SendMsg(d)
			fmt.Println("ReqName", info.Name)
		case 3:
			d, _ := json.Marshal(&Data{
				Code: data.Code,
				Data: "ws server test success",
			})
			_ = w.SendMsg(d)
		}
	}
}

func wsClose(w *WsConnection) {
	w.close()
}
