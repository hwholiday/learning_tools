package job_worker_mode

var JobQueue chan Goods

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
