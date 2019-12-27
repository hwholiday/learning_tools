package v1_server

import (
	"context"
	"fmt"
)

func (s service) Test(_ context.Context, req string) (res string, err error) {
	return fmt.Sprintf("%s_>>>>>>res", req), nil
}
