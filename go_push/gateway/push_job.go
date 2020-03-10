package gateway

type PushJob struct {
	Type   int
	roomId string
	info   []byte
}

type PushTask struct {
	JobChan          []chan *PushJob
	DistributionTask chan *PushJob
}

func NewPushTask() *PushTask {
	var p *PushTask
	p.DistributionTask = make(chan *PushJob, 10)
	p.JobChan = make([]chan *PushJob, len(roomTitle)) //一个房间对应一个job
	for i, _ := range roomTitle {
		go p.distributionTask(i) //分发任务
	}
	return p
}

func (p *PushTask) distributionTask(roomId int) {

}

func (p *PushTask) pushWork(roomId int) {
     var (
	 )
}
