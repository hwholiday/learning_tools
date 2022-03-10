package hystrix

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
	"testing"
	"time"
)

func TestHystrix(t *testing.T) {
	config := hystrix.CommandConfig{
		Timeout:                2000, //执行command的超时时间
		MaxConcurrentRequests:  8,    //command的最大并发量
		SleepWindow:            2000, //过多长时间，熔断器再次检测是否开启。单位毫秒
		ErrorPercentThreshold:  30,   //错误率 请求数量大于等于RequestVolumeThreshold并且错误率到达这个百分比后就会启动
		RequestVolumeThreshold: 5,    //请求阈值(一个统计窗口10秒内请求数量)  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
	}
	hystrix.ConfigureCommand("test", config)
	cbs, _, _ := hystrix.GetCircuit("test")
	defer hystrix.Flush()
	i := 1
	for {
		start1 := time.Now()
		hystrix.Do("test", RunFunc, func(e error) error {
			fmt.Println("服务器错误 触发 fallbackFunc 调用函数执行失败 : ", e)
			return nil
		})
		fmt.Println("请求次数:", i+1, ";用时:", time.Now().Sub(start1), ";熔断器开启状态:", cbs.IsOpen(), "请求是否允许：", cbs.AllowRequest())
		time.Sleep(time.Second)
		i++
	}
}

func RunFunc() error {
	rand.Seed(time.Now().Unix())
	if rand.Intn(10) > 5 {
		fmt.Println("[RunFunc] 执行失败")
		return errors.New("RunFunc ERROR")
	}
	fmt.Println("[RunFunc] 执行成功")
	return nil
}
