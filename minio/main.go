package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"net/url"
	"os"
	"time"
)

func main() {


}


func CheckErr(msg string, err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("%s :%v", msg, err.Error()))
		os.Exit(1)
	}
}
