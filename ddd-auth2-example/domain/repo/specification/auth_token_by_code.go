package specification

import (
	"context"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/repo"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
)

type AuthTokenByOpenId struct {
	OpenId string `json:"code"`
}

func NewAuthTokenSpecificationByoOenId(openId string) repo.AuthTokenSpecificationRepo {
	return &AuthTokenByOpenId{OpenId: openId}
}

func (m AuthTokenByOpenId) ParameterCheck(ctx context.Context) error {
	if m.OpenId == "" {
		return hcode.SysParameterErr
	}
	return nil
}

func (m AuthTokenByOpenId) ToSql(ctx context.Context) interface{} {
	return m.OpenId
}
