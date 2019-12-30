package v1_service

type Service interface {
}

type NewService func(Service) Service

type baseServer struct {
}

func New()  {
	
}
