package aggregate

import (
	"context"
	"github.com/google/uuid"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/dto"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/entity"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/obj"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/repo"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/tool"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type AuthCode struct {
	authCodeRepo repo.AuthCodeRepo
	data         dto.AuthCodeReq
	merchant     *entity.Merchant
}

func (a *AuthCode) CreateCode(ctx context.Context) (code string, err error) {
	var (
		openId string
		host   string
	)
	if host, err = a.data.GetRedirectUriHost(); err != nil {
		log.GetLogger().Error("[aggregate] AuthCode CreateCode GetRedirectUriHost", zap.Any("data", a.data), zap.Error(err))
		return "", hcode.ParameterErr
	}
	if host != a.merchant.Host {
		log.GetLogger().Error("[aggregate] AuthCode CreateCode host != a.merchant.Host", zap.Any("merchant.Host", a.merchant.Host), zap.Any("host", host), zap.Any("data", a.data))
		return "", hcode.ParameterErr
	}
	if !strings.Contains(a.merchant.Scope, a.data.Scope) {
		log.GetLogger().Error("[aggregate] AuthCode CreateCode Scope", zap.Any("merchant.Scope", a.merchant.Scope), zap.Any("data", a.data))
		return "", hcode.ParameterErr
	}
	openId, err = tool.AesECBEncrypt(a.data.APPID, []byte(strconv.Itoa(a.data.UID)))
	if err != nil {
		log.GetLogger().Error("[aggregate] AuthCode CreateCode AesECBEncrypt", zap.Any("data", a.data), zap.Error(err))
		return
	}
	code = strings.ReplaceAll(uuid.New().String(), "-", "")
	err = a.authCodeRepo.CreateCode(ctx, obj.CodeOpenId{
		Code:   code,
		OpenID: openId,
		APPID:  a.data.APPID,
		Scope:  a.data.Scope,
	})
	return
}
