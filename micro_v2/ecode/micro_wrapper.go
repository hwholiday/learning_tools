package ecode

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/server"
	errors2 "github.com/pkg/errors"
)

func ServerEcodeWrapper(svrName string) server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			err := fn(ctx, req, rsp)
			if code, ok := IsCode(errors2.Cause(err)); ok {
				return errors.New(svrName, code.Error(), int32(code.Code()))
			}
			return err
		}
	}
}

func ClientEcodeCallWrapper(fn client.CallFunc) client.CallFunc {
	return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
		err := fn(ctx, node, req, rsp, opts)
		if err != nil {
			if verr, ok := err.(*errors.Error); ok {
				if verr.Code == 408 {
					return Deadline
				}
				return Code(verr.Code)
			}
			return err
		}
		return nil
	}
}
