package dto

type UserSimple struct {
	OpenId   string `json:"open_id"`   // 用户唯一标示
	Username string `json:"user_name"` // 用户名
	Phone    string `json:"phone"`     // 手机号码
	Avatar   string `json:"avatar"`    // 头像
}
