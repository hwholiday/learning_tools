package src

import "fmt"

type Howie struct {
	Addr string
	Data string
}

func NewHowie()*Howie  {
	var d =new(Howie)
	d.Addr="127.0.0.1"
	d.Data="我是测试信息i"
	return d
}

func (h *Howie) Name() string  {
	fmt.Println("调用Name")
	return h.Data
}

func (h *Howie) Run()  {
	fmt.Println("调用RUN")
}


