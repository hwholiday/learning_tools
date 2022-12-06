package main

import "fmt"

type TaskA struct {
	BaseHandler
}

func (t *TaskA) Do(task *Task) error {
	task.TaskA = "TaskA"
	fmt.Println("TaskA")
	return nil
}

type TaskB struct {
	BaseHandler
}

func (t *TaskB) Do(task *Task) error {
	task.TaskB = "TaskB"
	fmt.Println("TaskB")
	return nil
}

type TaskC struct {
	BaseHandler
}

func (t *TaskC) Do(task *Task) error {
	task.TaskC = "TaskC"
	fmt.Println("TaskC")
	return nil
}
