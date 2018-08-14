package main

import "fmt"

func main(){
	Add()
}

func Add(agrs ...string)  {
	if len(agrs)<=0{
		fmt.Println("111")
	}
}