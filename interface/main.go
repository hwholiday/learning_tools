package main

import (
	"fmt"
	"learning_tools/interface/src"
)

type AA struct {
	a src.Agent
}

func (a *AA) TestAA() {
	fmt.Println("测试AA")
}
func main() {
	var A = &AA{}
	A.a = src.NewHowie()
	A.TestAA()
	A.a.Name()
	A.a.Run()
}
