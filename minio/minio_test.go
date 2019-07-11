package main

import (
	"bytes"
	"fmt"
	"github.com/minio/minio-go"
	"io/ioutil"
	"testing"
)

var minioClient *minio.Client

func init() {
	var err error
	minioClient, err = minio.New("127.0.0.1:9000", "AKIAINCHU2DIYAQ66TPA", "ZaYtz0d61fJXQ7djyXBX4yZ5ob8Kj/WNXtw6PJob", false)
	if err != nil {
		panic(err)
	}
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

func TestNewMultipartUpload(t *testing.T) {
	daya := minio.Core{Client: minioClient}
	v1, err := daya.NewMultipartUpload("test", "2.log", minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}
	//e35dd4ed-3c44-4567-90af-93aaedc773b0
	fmt.Println(v1)
}

func TestNewPutObjectPart(t *testing.T) {
	abuf := bytes.Repeat([]byte("aaaaa"), 1024*1024)//至少要5MB 才能使用分片上传
	bbuf := bytes.Repeat([]byte("b"), 1024*1024)
	daya := minio.Core{Client: minioClient}
	bucket:="test"
	name :="3.log"
	v1, err := daya.NewMultipartUpload(bucket, name, minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(v1)
	var completeParts []minio.CompletePart
	part, err := daya.PutObjectPart(bucket, name, v1, 1, bytes.NewReader(abuf), int64(len(abuf)), "", "", nil)
	if err != nil {
		panic(err)	}
	fmt.Println("1",part)
	completeParts = append(completeParts, minio.CompletePart{PartNumber: part.PartNumber, ETag: part.ETag})

	part, err = daya.PutObjectPart(bucket, name, v1, 2, bytes.NewReader(bbuf), int64(len(bbuf)), "", "", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("2",part.ETag)
	fmt.Println("2",part.PartNumber)
	fmt.Println("2",part.LastModified)
	fmt.Println("2",part.Size)
	completeParts = append(completeParts, minio.CompletePart{PartNumber: part.PartNumber, ETag: part.ETag})
	_, err = daya.CompleteMultipartUpload(bucket, name, v1, completeParts)
	if err != nil {
		panic(err)
	}
}
