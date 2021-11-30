package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/obj"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/repo"
	consts "github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/conf"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"go.uber.org/zap"
	"time"
)

var _ repo.AuthCodeRepo = (*AuthCode)(nil)

type AuthCode struct {
	repository
}

func (a *AuthCode) getCacheKey(data string) string {
	return fmt.Sprintf("%s%s", consts.AuthCodeCacheKey, data)
}

func (a *AuthCode) CreateCode(ctx context.Context, data obj.CodeOpenId) error {
	saveData, err := Marshal(&data)
	if err != nil {
		log.GetLogger().Error("[AuthToken] CreateAuthToken Marshal", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	if err := a.rds.Set(a.getCacheKey(data.Code), saveData, time.Second*60*2).Err(); err != nil {
		log.GetLogger().Error("[AuthCode] CreateCode Set", zap.Any("req", data), zap.Error(err))
		return hcode.RedisExecErr
	}
	return nil
}

func (a *AuthCode) DelCode(ctx context.Context, repo repo.AuthCodeSpecificationRepo) error {
	if err := a.rds.Del(a.getCacheKey(fmt.Sprint(repo.ToSql(ctx)))).Err(); err != nil {
		log.GetLogger().Error("[AuthCode] DelCode Del", zap.Any("req", repo.ToSql(ctx)), zap.Error(err))
		return hcode.RedisExecErr
	}
	return nil
}

func (a *AuthCode) QueryCode(ctx context.Context, repo repo.AuthCodeSpecificationRepo) (obj.CodeOpenId, error) {
	saveData, err := a.rds.Get(a.getCacheKey(fmt.Sprint(repo.ToSql(ctx)))).Bytes()
	if err != nil {
		if err == redis.Nil {
			return obj.CodeOpenId{}, hcode.RedisExecErr
		}
		log.GetLogger().Error("[AuthCode] QueryCode Get", zap.Any("req", repo.ToSql(ctx)), zap.Error(err))
		return obj.CodeOpenId{}, hcode.RedisExecErr
	}
	var codeOpenId obj.CodeOpenId
	if err := Unmarshal(saveData, &codeOpenId); err != nil {
		log.GetLogger().Error("[QueryAuthToken] Unmarshal", zap.Any("req", repo.ToSql(ctx)), zap.Error(err))
		return obj.CodeOpenId{}, hcode.RedisExecErr
	}
	return codeOpenId, nil
}
