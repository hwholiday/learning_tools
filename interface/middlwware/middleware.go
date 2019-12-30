package middlwware

import (
	"fmt"
	"time"
)

type ServiceMiddleware func(Service) Service

type logServer struct {
	next Service
	log  string
}

func LogMiddleware(name string) ServiceMiddleware {
	return func(next Service) Service {
		return logServer{
			next: next,
			log:  name,
		}
	}
}

func (mw logServer) Add(a, b int) (res int) {
	defer func(start time.Time) {
		fmt.Println("log", mw.log, "a > ", a, "b > ", b, "res > ", res, "time ", time.Since(start))
	}(time.Now())
	return mw.next.Add(a, b)
}
