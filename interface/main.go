package main

import (
	"test/interface/src"
	"fmt"
)

func main()  {
	var a src.Agent=src.NewHowie()
	a.Run()
	fmt.Println(a.Name())
}
