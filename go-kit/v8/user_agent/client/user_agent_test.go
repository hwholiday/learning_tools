package client

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	"github.com/hwholiday/learning_tools/go-kit/v8/user_agent/pb"
	"github.com/hwholiday/learning_tools/go-kit/v8/user_agent/src"
	"github.com/hwholiday/learning_tools/go-kit/v8/utils"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"testing"
	"time"
)

func TestNewUserAgentClient(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	client, err := NewUserAgentClient([]string{"127.0.0.1:2379"}, logger)
	if err != nil {
		t.Error(err)
		return
	}
	hy := utils.NewHystrix("调用错误服务降级")
	cbs, _, _ := hystrix.GetCircuit("login")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 100)
		userAgent, err := client.UserAgentClient()
		if err != nil {
			t.Error(err)
			return
		}
		err = hy.Run("login", func() error {
			_, err := userAgent.Login(context.Background(), &pb.Login{
				Account:  "hwholiday",
				Password: "123456",
			})
			if err != nil {
				return err
			}
			//fmt.Println(ack.Token)
			return nil
		})
		fmt.Println("熔断器开启状态:", cbs.IsOpen(), "请求是否允许：", cbs.AllowRequest())
		if err != nil {
			t.Log(err)
		}
	}

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
	md := metadata.Pairs(src.ContextReqUUid, UUID)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	for i := 0; i < 20; i++ {
		res, err := userClient.RpcUserLogin(ctx, &pb.Login{
			Account:  "hwholiday",
			Password: "123456",
		})
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(res.Token)
	}
}
