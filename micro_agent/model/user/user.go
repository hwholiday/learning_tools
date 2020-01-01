package user

import (
	"fmt"
	user_agent "micro_agent/proto/user"
	"sync"
)
var (
	s *service
	m sync.RWMutex
)

type service struct {}

type Service interface {
	UserInfo(msg *user_agent.ReqMsg)(info string,err error)
}

func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	s = &service{}
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}






