package main

import "fmt"

type Msg string

type Msg2 struct {
	m Msg
}

type Msg3 struct {
	m Msg2
}

func NewMsg() Msg {
	return Msg("123")
}

func NewMsg2(m Msg) Msg2 {
	return Msg2{
		m: m,
	}
}

func NewMsg3(m Msg2) Msg3 {
	return Msg3{
		m: m,
	}
}
func (m Msg3) Start() {
	fmt.Println("111")
}
