package obj

type CodeOpenId struct {
	Code   string `json:"code"`
	OpenID string `json:"open_id"`
	APPID  string `json:"appid"`
	Scope  string `json:"scope"`
}
