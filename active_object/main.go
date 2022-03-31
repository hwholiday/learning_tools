package main

import (
	"fmt"
	"time"
)

//Go并发设计模式之 Active Object
//https://colobu.com/2019/07/02/concurrency-design-patterns-active-object/
func main() {
	info := NewService(1)
	for i := 0; i < 20; i++ {
		info.Add()
		fmt.Println(info.val) //不是同步返回值
		//info.Lessen()
		//fmt.Println(info.val)
	}
	fmt.Println(info.val)
	time.Sleep(time.Hour)
}
