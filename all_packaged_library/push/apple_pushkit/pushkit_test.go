package apple_pushkit

import (
	"testing"
	"fmt"
)

func TestInitPushKit(t *testing.T) {
	push,err:=InitPushKit("./131231P.p12","pwd",true)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(push.Push("123123",[]byte(`{"newsid":{"content":"test",},"badge":{"badge":"0"}}`)))
}