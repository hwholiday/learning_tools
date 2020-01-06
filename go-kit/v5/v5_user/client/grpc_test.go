package client

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"learning_tools/all_packaged_library/logtool"
	"learning_tools/go-kit/v5/v5_user/pb"
	"testing"
)

func TestGrpcClient(t *testing.T) {
	logger := logtool.NewLogger(
		logtool.SetAppName("go-kit"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
	conn, err := grpc.DialContext(context.Background(), ":8881", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	svr := NewGRPCClient(conn, logger)
	ack, err := svr.Login(context.Background(), &pb.Login{
		Account:  "hw",
		Password: "123",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ack.Token)
}

func TestGrpc(t *testing.T) {

	serviceAddress := "127.0.0.1:8881"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	userClient := pb.NewUserClient(conn)
	res, err := userClient.RpcUserLogin(context.Background(), &pb.Login{
		Account:  "hw",
		Password: "123",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res.Token)

}
