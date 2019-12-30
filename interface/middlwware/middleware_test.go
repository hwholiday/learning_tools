package middlwware

import (
	"testing"
)

func Test_middleware(t *testing.T) {
	svc := NewService("")
	t.Log(svc.Add(1, 2))
}
