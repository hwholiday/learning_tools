package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Message struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data []byte `json:"data,omitempty"` //json.RawMessage
}

type User struct {
	Name string `json:"name,omitempty"`
	Pwd  string `json:"pwd,omitempty"`
}

func TestJsonV1(t *testing.T) {
	var user User
	user.Name = "test_user"
	user.Pwd = "test_pwd"
	var ms Message
	ms.Code = 200
	ms.Msg = "success"
	userJson, _ := json.Marshal(&user)
	ms.Data = userJson
	msData, _ := json.Marshal(ms)
	t.Log(string(msData))
	fmt.Println(string(msData))

}

func Test_json(t *testing.T) {
	//	var data="{\"code\":200,\"msg\":\"success\",\"data\":{\"name\":\"test_user\",\"pwd\":\"test_pwd\"}}" //json.RawMessage
	var data = "{\"code\":200,\"msg\":\"success\",\"data\":\"eyJuYW1lIjoidGVzdF91c2VyIiwicHdkIjoidGVzdF9wd2QifQ==\"}"
	var ms Message
	var user User
	_ = json.Unmarshal([]byte(data), &ms)
	fmt.Println(ms.Code)
	fmt.Println(ms.Msg)
	_ = json.Unmarshal(ms.Data, &user)
	fmt.Println(user.Name)
	fmt.Println(user.Pwd)
}
