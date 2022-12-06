package main

import "fmt"

type Task struct {
	Name  string
	TaskA string
	TaskB string
	TaskC string
}

type TaskHandler interface {
	NextHandler(TaskHandler) TaskHandler
	Do(*Task) error
	Execute(*Task) error
}

type BaseHandler struct {
	nextTaskHandler TaskHandler
}

func (h *BaseHandler) NextHandler(task TaskHandler) TaskHandler {
	h.nextTaskHandler = task
	return task
}

func (h *BaseHandler) Execute(task *Task) error {
	if h.nextTaskHandler != nil {
		if err := h.nextTaskHandler.Do(task); err != nil {
			return err
		}
		return h.nextTaskHandler.Execute(task)
	}
	return nil
}

type StartHandler struct {
	BaseHandler
}

func (s *StartHandler) Do(task *Task) error {
	fmt.Println("StartHandler")
	return nil
}
