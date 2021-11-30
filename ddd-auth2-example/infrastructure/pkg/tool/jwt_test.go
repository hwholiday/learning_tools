package tool

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateAuthToken(t *testing.T) {
	res, err := CreateAuthToken(JwtTokenData{
		OpenId: "aaaa",
		AppId:  "111",
	}, time.Second*3)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", res)
	time.Sleep(time.Second * 5)
	fmt.Println(CheckAuthToken(res.Token))
}
