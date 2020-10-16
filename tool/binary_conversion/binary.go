package main

import "fmt"

func main() {
	var switchs uint64
	fmt.Printf("%064b\n", switchs)
	switchs = 1<<63 | 1<<2 | 1<<3 //直接替换
	fmt.Printf("%064b\n", switchs)
	switchs = 1 << 0
	fmt.Printf("%064b\n", switchs)
	switchs = 1<<62 | 1<<61 | 1<<60 //修改某一些的东西
	fmt.Printf("%064b\n", switchs)
	fmt.Println(switchs)
	fmt.Printf("%064b\n", 1<<62)
	fmt.Printf("%064b\n", 8070450532247928832)
	fmt.Println(1 << 2 & 8070450532247928832)  //判断第二位是不是0
	fmt.Println(1 << 62 & 8070450532247928832) //判断第二位是不是0
	fmt.Println(1 << 1 & 8070450532247928832)  //判断第二位是不是0
}
