package main

import "fmt"

type obj1 struct {
}

func (c *obj1)Add()  {
	fmt.Println("obj1")
}
type obj2 struct {
	obj obj1
}

func (c *obj2)Add()  {
	fmt.Println("obj2")
}

func main() {
	var i obj2
	i.obj.Add()

	var a=[]string{"1","2","3"}
	dd1(a)
	fmt.Println("--------------")
	fmt.Println(a)
}

func dd1(a []string)  {
	a=append(a,"5")
	fmt.Println("22222222222222222222222222222222222")
	fmt.Println(a)
}
