package main

import (
	"fmt"
	"os"
)

//分片上传文件
//分片文件小于5M 再redis缓存到5M在上传
func main() {


}

func CheckErr(msg string, err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("%s :%v", msg, err.Error()))
		os.Exit(1)
	}
}
