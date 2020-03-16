package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	"go.uber.org/zap"
	"learning_tools/all_packaged_library/logtool"
	"learning_tools/go-kit/v11/user_agent/src"
	"learning_tools/go-kit/v11/utils"
	"net/http"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	zapLogger := logtool.NewLogger(
		logtool.SetAppName("go-kit-v11-client"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
	utils.NewLoggerServer()
	client, err := NewUserAgentClient([]string{"127.0.0.1:2379"}, logger)
	if err != nil {
		utils.GetLogger().Debug("[NewUserAgentClient]", zap.Error(err))
		return
	}
	hy := utils.NewHystrix("调用错误服务降级")
	cbs, _, _ := hystrix.GetCircuit("login")
	//curl  http://0.0.0.0:10010/login -X POST -d "account=hwholiday&password=123456"
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		account := request.Form.Get("account")
		password := request.Form.Get("password")
		userAgent, err := client.UserAgentClient()
		if err != nil {
			utils.GetLogger().Debug("[UserAgentClient]", zap.Error(err))
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		err = hy.Run("login", func() error {
			ack, err := userAgent.Login(context.Background(), src.Login{
				Account:  account,
				Password: password,
			})
			if err != nil {
				zapLogger.Error("[login]", zap.Error(err))
				_, _ = writer.Write([]byte(err.Error()))
				return err
			}
			_, _ = writer.Write([]byte(ack.Token))
			zapLogger.Error("[login]", zap.Any("返回值", ack.Token))
			return nil
		})
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
		}
		zapLogger.Debug("熔断器", zap.Any("开启状态", cbs.IsOpen()), zap.Any("请求是否允许：", cbs.AllowRequest()))
	})
	fmt.Println("服务启动成功 监听端口 10010")
	er := http.ListenAndServe("0.0.0.0:10010", nil)
	if er != nil {
		fmt.Println("ListenAndServe: ", er)
	}
}
