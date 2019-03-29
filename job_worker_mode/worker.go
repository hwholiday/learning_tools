package job_worker_mode

var JobQueue chan Goods

type Worker struct {
	WorkerPool chan chan Goods
	JobChannel chan Goods
	Quit       chan bool
}

