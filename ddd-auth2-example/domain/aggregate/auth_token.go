package aggregate

import (
	"context"
	"fmt"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/dto"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/domain/repo"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/hcode"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/tool"
	"strconv"
)

type AuthToken struct {
	openId        string
	appId         string
	authTokenRepo repo.AuthTokenRepo
}

func (a AuthToken) GetUserInfo(ctx context.Context) (userSimple dto.UserSimple, err error) {
	var (
		uid int
	)
	uidByte, err := tool.AesECBDecrypt(fmt.Sprint(a.appId), a.openId)
	if err != nil {
		err = hcode.ServerErr
		return
	}
	uid, err = strconv.Atoi(string(uidByte))
	if err != nil {
		err = hcode.TranErr
		return
	}
	fmt.Println("uid", uid)
	// TODO 这里可以在 adpter 层去实现获取用户信息
	return dto.UserSimple{
		OpenId:   a.openId,
		Username: "",
		Phone:    "",
		Avatar:   "",
	}, nil
}
