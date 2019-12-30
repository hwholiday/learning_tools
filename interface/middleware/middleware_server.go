package middleware


type Service interface {
	Add(a, b int) int
}

type baseServer struct{}

func NewBaseServer() Service {
	return baseServer{}
}

func (s baseServer) Add(a, b int) int {
	return a + b
}

func NewService(s string) Service {
	var svc Service
	{
		svc = NewBaseServer()
		svc = LogMiddleware(s)(svc)
		svc = LogV2Middleware(s)(svc)
	}
	return svc
}
