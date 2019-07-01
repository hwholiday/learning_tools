package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := strconv.FormatUint(212121212, 36) //10进制转36进制
	fmt.Println(data)
	fmt.Println(strconv.ParseUint(data, 36, 64)) //36进制转10进制
}
