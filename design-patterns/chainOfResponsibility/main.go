package main

import "fmt"

func main() {
	s := StartHandler{}
	s.NextHandler(&TaskA{}).NextHandler(&TaskB{}).NextHandler(&TaskC{})
	var task = &Task{
		Name: "test",
	}
	if err := s.Execute(task); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", task)
}
