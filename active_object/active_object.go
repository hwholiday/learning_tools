package main

import (
	"fmt"
	"time"
)

const (
	Add    = 1
	Lessen = 2
)

type Service struct {
	queue chan int
	val   int
}

func NewService(buffer int) *Service {
	info := &Service{queue: make(chan int, buffer)}
	go info.schedule()
	return info
}

func (s *Service) schedule() {
	for {
		select {
		case i := <-s.queue:
			if i == Add {
				fmt.Println("s.val++")
				s.val++
			} else if i == Lessen {
				fmt.Println("s.val--")
				s.val--
			}
		case <-time.After(time.Second * 5):
			fmt.Println("5秒没有写入")
		}
	}
}

func (s *Service) Add() {
	fmt.Println("s.queue <- Add")
	s.queue <- Add
}

func (s *Service) Lessen() {
	fmt.Println("s.queue <- Lessen")
	s.queue <- Lessen
}
