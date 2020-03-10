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
	GetRoomManage().AllRoom.Range(func(key, value interface{}) bool {
		return true
	})
	return p
}

func (p *PushTask) distributionTask() {

}
