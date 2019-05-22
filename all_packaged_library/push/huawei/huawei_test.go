package huawei

import (
	"testing"
	"fmt"
	"time"
)

func TestHuaweiPush_GetToken(t *testing.T) {
	push := NewHuaweiPush("https://login.cloud.huawei.com/oauth2/v2/token", "100358845", "bee1d8f704b1bc278bea7f5427cb0f8a", "https://api.push.hicloud.com/pushsend.do", true)
	var in ReqPush
	in.DeviceTokenList = []string{"0862791036717594300002894200CN01"}
	in.Ver = "1"
	in.NspTs = "1545113076"
	in.Payload = `{"hps":{"msg":{"type":1,"body":{"key":"value"}}}} `
	for {
		time.Sleep(time.Second * 3)
		//判断ResPush中的code是不是等于80000000可以测试是否成功
		//{"requestId":"154536200334112580200","msg":"Success","code":"80000000"}
		fmt.Println(push.Push(&in))
	}
}
