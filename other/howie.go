package main

import "fmt"

func main1(){
	Add()
}

func Add(agrs ...string)  {
	if len(agrs)<=0{
		fmt.Println("111")
	}
}