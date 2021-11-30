package specification

import (
	"context"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/repo"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
)

type AuthCodeByCode struct {
	Code string `json:"code"`
}

func NewAuthCodeSpecificationByCode(code string) repo.AuthCodeSpecificationRepo {
	return &AuthCodeByCode{Code: code}
}

func (m AuthCodeByCode) ParameterCheck(ctx context.Context) error {
	if m.Code == "" {
		return hcode.SysParameterErr
	}
	return nil
}

func (m AuthCodeByCode) ToSql(ctx context.Context) interface{} {
	return m.Code
}
