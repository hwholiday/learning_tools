package ecode

var (
	OK = add(0) // 正确

	RequestErr         = add(-400) // 请求错误
	Unauthorized       = add(-401) // 未认证
	AccessDenied       = add(-403) // 访问权限不足
	NothingFound       = add(-404) // 资源不存在
	MethodNotAllowed   = add(-405) // 不支持该方法
	ParamErr           = add(-406) // 参数错误
	Canceled           = add(-498) // 客户端取消请求
	ServerErr          = add(-500) // 服务器错误
	ServiceUnavailable = add(-503) // 过载保护,服务暂不可用
	Deadline           = add(-504) // 服务调用超时
	LimitExceed        = add(-509) // 超出限制
)

func init() {
	cms := map[int]map[string]string{
		ParamErr.Code(): {
			"cn": "你好世界",
			"en": "hello world",
		},
	}
	Register(cms)
}
