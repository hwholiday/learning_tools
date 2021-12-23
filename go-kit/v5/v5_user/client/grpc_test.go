package client

import (
	"context"
	"github.com/hwholiday/learning_tools/all_packaged_library/logtool"
	"github.com/hwholiday/learning_tools/go-kit/v5/v5_user/pb"
	"github.com/hwholiday/learning_tools/go-kit/v5/v5_user/v5_service"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestGrpcClient(t *testing.T) {
	logger := logtool.NewLogger(
		logtool.SetAppName("go-kit"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
	conn, err := grpc.Dial("127.0.0.1:8881", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	svr := NewGRPCClient(conn, logger)
	ack, err := svr.Login(context.Background(), &pb.Login{
		Account:  "hwholiday",
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ack)
}

func TestGrpc(t *testing.T) {
	serviceAddress := "127.0.0.1:8881"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	userClient := pb.NewUserClient(conn)
	UUID := uuid.NewV5(uuid.NewV4(), "req_uuid").String()
	md := metadata.Pairs(v5_service.ContextReqUUid, UUID)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := userClient.RpcUserLogin(ctx, &pb.Login{
		Account:  "hw",
		Password: "123",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res.Token)

}
