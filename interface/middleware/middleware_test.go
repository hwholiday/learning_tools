package middleware

import (
	"testing"
)

func Test_middleware(t *testing.T) {
	t.Log("开始运行")
	svc := NewService("")
	t.Log("结果", svc.Add(1, 2))
}
