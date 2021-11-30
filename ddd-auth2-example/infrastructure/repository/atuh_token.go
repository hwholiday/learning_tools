package repository

import (
	"context"
	"fmt"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/obj"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/repo"
	consts "github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/conf"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"go.uber.org/zap"
)

var _ repo.AuthTokenRepo = (*AuthToken)(nil)

type AuthToken struct {
	repository
}

func (a *AuthToken) getCacheKey(data string) string {
	return fmt.Sprintf("%s%s", consts.AuthTokenCacheKey, data)
}

func (a *AuthToken) CreateAuthToken(ctx context.Context, data obj.AuthToken) error {
	saveData, err := Marshal(&data)
	if err != nil {
		log.GetLogger().Error("[AuthToken] CreateAuthToken Marshal", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	if err := a.rds.Set(a.getCacheKey(data.OpenID), saveData, consts.AuthRefreshTokenCacheKeyTimeout).Err(); err != nil {
		log.GetLogger().Error("[AuthToken] CreateAuthToken Set", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	return nil
}

func (a *AuthToken) UpdateAuthToken(ctx context.Context, data obj.AuthToken) error {
	saveData, err := Marshal(&data)
	if err != nil {
		log.GetLogger().Error("[AuthToken] UpdateAuthToken Marshal", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	if err := a.rds.Set(a.getCacheKey(data.OpenID), saveData, consts.AuthRefreshTokenCacheKeyTimeout).Err(); err != nil {
		log.GetLogger().Error("[AuthToken] UpdateAuthToken Set", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	return nil
}

func (a *AuthToken) QueryAuthToken(ctx context.Context, repo repo.AuthTokenSpecificationRepo) (obj.AuthToken, error) {
	data, err := a.rds.Get(a.getCacheKey(fmt.Sprint(repo.ToSql(ctx)))).Bytes()
	if err != nil {
		log.GetLogger().Error("[QueryAuthToken] Get", zap.Any("req", repo.ToSql(ctx)), zap.Error(err))
		return obj.AuthToken{}, hcode.RedisExecErr
	}
	var authToken obj.AuthToken
	if err := Unmarshal(data, &authToken); err != nil {
		log.GetLogger().Error("[QueryAuthToken] Unmarshal", zap.Any("req", repo.ToSql(ctx)), zap.Error(err))
		return obj.AuthToken{}, hcode.RedisExecErr
	}
	return authToken, nil
}
