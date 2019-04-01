package main

type DisPatcher struct {
	Worker    chan chan Goods
	JobQueue  chan Goods
	MaxWorker int
}

func NewDisPatcher(maxWorker, maxQueue int) *DisPatcher {
	return &DisPatcher{
		Worker: make(chan chan Goods, maxWorker), MaxWorker: maxWorker,
		JobQueue: make(chan Goods, maxQueue),
	}
}

func (d *DisPatcher) Run() {
	for i := 0; i < d.MaxWorker; i++ {
		worker := NewWorker(d.Worker)
		worker.Start()
	}
	go d.dispatch()
}

func (d *DisPatcher) dispatch() {
	for {
		select {
		case goods := <-d.JobQueue:
			go func(good Goods) {
				jobChannel := <-d.Worker
				jobChannel <- goods
			}(goods)
		}
	}
}
