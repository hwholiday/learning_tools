package main

import "fmt"

func main() {
	var switchs uint64
	fmt.Printf("%064b\n", switchs)
	switchs = 1<<63 | 1<<2 | 1<<3 //直接替换
	fmt.Printf("%064b\n", switchs)
	switchs += 1<<4 | 1<<5 | 1<<6 //修改某一些的东西
	fmt.Printf("%064b\n", switchs)
	fmt.Println(1 << 10 & switchs) //判断第二位是不是0
}
