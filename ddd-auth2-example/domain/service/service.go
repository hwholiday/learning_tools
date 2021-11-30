package service

import (
	"context"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/aggregate"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/dto"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/repository"
)

type AuthSrv interface {
	CreateCodeOpenId(ctx context.Context, req dto.AuthCodeReq) (string, error)

	CreateToken(ctx context.Context, data dto.ProduceAuthTokenReq) (authTokenSimple dto.AuthTokenSimple, err error)
	RefreshAccessToken(ctx context.Context, data dto.RefreshAccessTokenReq) (authTokenSimple dto.AuthTokenSimple, err error)
	GetUserInfo(ctx context.Context, data dto.OpenIdTokenReq) (user dto.UserSimple, err error)
	CheckToken(ctx context.Context, data dto.OpenIdTokenReq) (err error)
}

type AuthService struct {
	*AuthCode
	*AuthToken
	*Merchant
}

func NewService(r *repository.Repository, a *aggregate.Factory) AuthSrv {
	s := &AuthService{
		AuthCode: &AuthCode{
			factory:      a.AuthFactory,
			authCodeRepo: r.AuthCode,
		},
		AuthToken: &AuthToken{
			factory:       a.AuthFactory,
			authTokenRepo: r.AuthToken,
		},
		Merchant: &Merchant{
			merchantRepo: r.Merchant,
		},
	}
	//fmt.Println(s.Merchant.CreateMerChant(context.Background(),&entity.Merchant{
	//	APPID:      "1234567899",
	//	Host:       "127.0.0.1:8888",
	//	Secret:     "123",
	//	Scope:      "snsapi_base,snsapi_userinfo",
	//	StartTime:  time.Now(),
	//	CreateTime: tool.GetTimeUnixMilli(),
	//	UpdateTime: tool.GetTimeUnixMilli(),
	//}))
	return s
}
