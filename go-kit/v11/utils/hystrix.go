package utils

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"sync"
)

var config = hystrix.CommandConfig{
	Timeout:                5000, //执行command的超时时间(毫秒)
	MaxConcurrentRequests:  8,    //command的最大并发量
	SleepWindow:            1000, //过多长时间，熔断器再次检测是否开启。单位毫秒
	ErrorPercentThreshold:  30,   //错误率 请求数量大于等于RequestVolumeThreshold并且错误率到达这个百分比后就会启动
	RequestVolumeThreshold: 5,    //请求阈值(一个统计窗口10秒内请求数量)  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
}

type runFunc func() error

type Hystrix struct {
	loadMap  *sync.Map
	fallback string
}

func NewHystrix(msg string) *Hystrix {
	return &Hystrix{
		loadMap:  new(sync.Map),
		fallback: msg,
	}
}

func (s *Hystrix) Run(name string, run runFunc) error {
	if _, ok := s.loadMap.Load(name); !ok {
		hystrix.ConfigureCommand(name, config)
		s.loadMap.Store(name, name)
	}
	err := hystrix.Do(name, func() error {
		return run()
	}, func(err error) error {
		fmt.Println("运行 run 方法错误 ", err)
		//return nil
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
