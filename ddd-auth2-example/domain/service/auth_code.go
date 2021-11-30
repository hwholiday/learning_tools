package service

import (
	"context"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/aggregate"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/dto"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/repo"
)

type AuthCode struct {
	factory      *aggregate.AuthFactory
	authCodeRepo repo.AuthCodeRepo
}

func (a *AuthCode) CreateCodeOpenId(ctx context.Context, req dto.AuthCodeReq) (string, error) {
	if err := req.Check(); err != nil {
		return "", err
	}
	f, err := a.factory.NewAuthCode(ctx, req)
	if err != nil {
		return "", err
	}
	return f.CreateCode(ctx)
}
