package huawei

type HuaweiPush struct {
	TokenUrl             string `json:"access_token_api"`
	AppId                string `json:"app_id"`
	AppSecret            string `json:"app_secret"`
	PushUrl              string `json:"push_url"`
	AutoCacheAccessToken bool   `json:"auto_cache_access_token"` //是否启用自动缓存token过期自动获取,设置为true可以不再调用getToken方法
}

type ResToken struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type ReqPush struct {
	Ver             string   `json:"ver"` //用来解决大版本升级的兼容问题;
	AccessToken     string   `json:"access_token"`
	NspTs           string   `json:"nsp_ts"`  //不传入使用默认值
	NspSvc          string   `json:"nsp_svc"` //不传入使用默认值
	DeviceTokenList []string `json:"device_token_list"`
	ExpireTime      string   `json:"expire_time"` //可选
	Payload         string   `json:"payload"`
}

type ResPush struct {
	RequestId string `json:"requestId"`
	Msg       string `json:"msg"`
	Code      string    `json:"code"`
	Error     string    `json:"error"`
}
