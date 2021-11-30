package entity

import (
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
	"time"
)

type Merchant struct {
	APPID      string    `json:"appid" bson:"appid"`
	Host       string    `json:"host" bson:"host"`
	Secret     string    `json:"secret" bson:"secret"`
	Scope      string    `json:"scope"` // 授权范围 ,号分割
	StartTime  time.Time `json:"start_time" bson:"start_time"`
	CreateTime int64     `json:"create_time" bson:"create_time"`
	UpdateTime int64     `json:"update_time" bson:"update_time"`
}

func (m *Merchant) CheckBase() error {
	if len(m.APPID) != 10 {
		return hcode.ParameterErr
	}
	if m.Host == "" || m.Secret == "" || m.Scope == "" {
		return hcode.ParameterErr
	}
	return nil
}
