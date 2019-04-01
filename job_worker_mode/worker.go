package main

import "fmt"



type Worker struct {
	WorkerPool chan chan Goods
	JobChannel chan Goods
	Quit       chan bool
}

func NewWorker(pool chan chan Goods) *Worker {
	return &Worker{
		WorkerPool: pool,
		JobChannel: make(chan Goods),
		Quit:       make(chan bool),
	}
}
func (w *Worker) Start() {
	go func() {
		for {

			//将当前工作者注册到工作队列中
			w.WorkerPool <- w.JobChannel
			select {
			case goods := <-w.JobChannel:
				//执行该程序
				goods.UpdateServer()
			case <-w.Quit:
				fmt.Println("worker服务停止")
				return
			}
		}
	}()
}
func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}
