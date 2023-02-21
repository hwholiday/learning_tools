package timewheel

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Job func(interface{})

type TimeWheel struct {
	interval     int64
	slots        []*list.List
	slotsNum     int64
	currentSlots int64
	ticker       *time.Ticker
	tasks        sync.Map
	addTaskCh    chan *Task
	removeTaskCh chan string
	closeCh      chan struct{}
}

type Task struct {
	ID         string
	createTime int64
	slots      int64
	circle     int64 // 多少圈
	job        Job
}

func NewTimeWheel() *TimeWheel {
	tw := &TimeWheel{
		interval:     1,
		slotsNum:     60 * 60 * 24,
		addTaskCh:    make(chan *Task),
		removeTaskCh: make(chan string),
		closeCh:      make(chan struct{}),
		ticker:       time.NewTicker(time.Second * time.Duration(1)),
	}
	tw.slots = make([]*list.List, tw.slotsNum)
	for i := int64(0); i < tw.slotsNum; i++ {
		tw.slots[i] = list.New()
	}
	tw.run()
	return tw
}

func (t *TimeWheel) AddFn(ID string, job Job, delay int64) {
	circle, slots := t.getCircleAndSlots(delay)
	task := &Task{
		ID:         ID,
		createTime: time.Now().Unix(),
		slots:      slots,
		circle:     circle,
		job:        job,
	}
	t.addTaskCh <- task
}
func (t *TimeWheel) run() {
	go func() {
		for {
			select {
			case _ = <-t.ticker.C:
				fmt.Println(time.Now().Format(time.DateTime))
				t.runTask()
			case task := <-t.addTaskCh:
				ele := t.slots[task.slots].PushBack(task)
				t.tasks.Store(task.ID, ele)
			case id := <-t.removeTaskCh:
				if val, ok := t.tasks.Load(id); ok {
					task := val.(*list.Element).Value.(*Task)
					ele := t.slots[task.slots].Remove(val.(*list.Element))
					t.tasks.Store(task.ID, ele)
				}
			case _ = <-t.closeCh:
				t.ticker.Stop()
				break
			}
		}
	}()
}

func (t *TimeWheel) runTask() {
	tasks := t.slots[t.currentSlots]
	if tasks != nil {
		for item := tasks.Front(); item != nil; item = item.Next() {
			task := item.Value.(*Task)
			if task.circle > 0 {
				task.circle--
				item = item.Next()
				continue
			}
			go task.job(task.ID)
			t.tasks.Delete(task.ID)
			tasks.Remove(item)
		}
	}
	if t.currentSlots == t.currentSlots-1 {
		t.currentSlots = 0
	} else {
		t.currentSlots++
	}
}

func (t *TimeWheel) getCircleAndSlots(delay int64) (circle, slots int64) {
	circle = delay / t.slotsNum
	slots = delay - (t.slotsNum * circle) + t.currentSlots
	if slots == t.currentSlots && circle > 0 {
		circle--
	}
	slots--
	return
}
