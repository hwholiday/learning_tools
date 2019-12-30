package middlwware

import (
	"testing"
)

func Test_middleware(t *testing.T) {
	svc := NewService("日志中间件")
	t.Log(svc.Add(1, 2))
}
