package v1_server

import "context"

type Service interface {
	Test(_ context.Context, req string) (res string, err error)
}

type NewService func(Service) Service

type service struct {
}
