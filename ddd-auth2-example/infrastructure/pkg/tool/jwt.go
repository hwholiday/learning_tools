package tool

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"go.uber.org/zap"
)

const (
	key = "1Wsaw2V3zx265x)!@#"
)

type JwtToken struct {
	Token         string
	TokenTimeline int64
}

type JwtTokenData struct {
	OpenId string
	AppId  string
	Scope  string
}

func CreateAuthToken(in JwtTokenData, t time.Duration) (JwtToken, error) {
	var err error
	data := JwtToken{
		TokenTimeline: time.Now().Add(t).Unix(),
	}
	data.Token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized":  true,
		"open_id":     in.OpenId,
		"app_id":      in.AppId,
		"scope":       in.Scope,
		"create_time": time.Now().Unix(),
		"exp":         data.TokenTimeline,
	}).SignedString([]byte(key))
	if err != nil {
		return data, err
	}
	return data, nil
}

func CheckAuthToken(authToken string) (out JwtTokenData, err error) {
	var parsed *jwt.Token
	parsed, err = jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		log.GetLogger().Error("CheckAuthToken", zap.Any("authToken", authToken), zap.Error(err))
		err = hcode.TokenValidErr
		return
	}
	if !parsed.Valid {
		err = hcode.TokenValidErr
		return
	}
	claims := parsed.Claims.(jwt.MapClaims)
	out.OpenId = claims["open_id"].(string)
	out.AppId = claims["app_id"].(string)
	out.Scope = claims["scope"].(string)
	return
}
