package handler

import (
	"context"
	"fmt"
	user_agent "micro_agent/proto/user"
)

func (s *Service)RpcUserInfo(ctx context.Context,req *user_agent.ReqMsg,res *user_agent.ResMsg)error  {
	fmt.Println(s.userServer.UserInfo(req))
	return nil
}