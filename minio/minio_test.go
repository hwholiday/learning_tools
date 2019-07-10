package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"io/ioutil"
	"testing"
)

var minioClient *minio.Client

func init() {
	var err error
	minioClient, err = minio.New("127.0.0.1:9000", "AKIAINCHU2DIYAQ66TPA", "ZaYtz0d61fJXQ7djyXBX4yZ5ob8Kj/WNXtw6PJob", false)
	CheckErr("minio.New", err)

}
func TestGetObject(t *testing.T) {
	var opt minio.GetObjectOptions
	err := opt.SetRange(1, 2)
	if err != nil {
		panic(err)
	}
	object, err := minioClient.GetObject("test", "1.txt", opt)
	if err != nil {
		panic(err)
	}
	defer object.Close()
	data, err := ioutil.ReadAll(object)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
