package model

import (
	"testing"
	"encoding/json"
	"fmt"
)

func Test(t *testing.T)  {
	data:=&Announcement{}
	d,_:=json.Marshal(data)
	fmt.Println(string(d))
}