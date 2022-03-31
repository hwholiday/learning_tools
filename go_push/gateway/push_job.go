package gateway

import "fmt"

type PushJob struct {
	Type     int
	PushType int
	RoomId   int
	Info     string
}

var pushTask *PushTask

type PushTask struct {
	JobChan          []chan *PushJob
	DistributionTask chan *PushJob
}

type PushManage interface {
	Push(job *PushJob)
}

func GetPushManage() PushManage {
	return pushTask
}

func NewPushTask(roomLen, workNum, taskNum int) {
	pushTask = &PushTask{
		JobChan:          make([]chan *PushJob, roomLen),
		DistributionTask: make(chan *PushJob, taskNum),
	}
	for i := 0; i < roomLen; i++ {
		//可以一个房间开多个pushWork(i)
		pushTask.JobChan[i] = make(chan *PushJob, roomLen)
		go pushTask.pushWork(i) //分发任务
	}
	for i := 0; i < workNum; i++ {
		go pushTask.distributionTask()
	}
}

func (p *PushTask) Push(job *PushJob) {
	p.DistributionTask <- job
}

func (p *PushTask) distributionTask() {
	var (
		pushJob *PushJob
	)
	for {
		select {
		case pushJob = <-p.DistributionTask:
			// 分发
			if pushJob.Type == 1 {
				GetRoomManage().PushAll(&WSMessage{
					Type: pushJob.PushType,
					Data: pushJob.Info,
				})
			} else if pushJob.Type == 2 {
				p.JobChan[pushJob.RoomId] <- pushJob
			}
		}
	}
}

func (p *PushTask) pushWork(roomId int) {
	var (
		err error
		job *PushJob
	)
	for {
		select {
		case job = <-p.JobChan[roomId]:
			if err = GetRoomManage().PushRoom(roomId, &WSMessage{
				Type: job.PushType,
				Data: job.Info,
			}); err != nil {
				fmt.Println(err)
			}
		}
	}
}
