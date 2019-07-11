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
	//s3.amazonaws.com
	//127.0.0.1:9000
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
func TestPutObject(t *testing.T)  {
	/*_,err:=minioClient.PutObject("test","3",strings.NewReader("3"),int64(len("1")),minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}*/
	info,err:=minioClient.StatObject("test","1",minio.StatObjectOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
	fmt.Println(info.ETag)
}


func TestComposeObject(t *testing.T)  {
    bucketName:="test"
	var srcs =make([]minio.SourceInfo,0)
	for i:=0;i<3;i++{
		name:=fmt.Sprintf("sheet_%d",i)
		bbuf := bytes.Repeat([]byte(fmt.Sprint(i+1)), 1024*1025*5)//至少要5M

		index,err:=minioClient.PutObject(bucketName,name,bytes.NewReader(bbuf),int64(len(bbuf)),minio.PutObjectOptions{})
        if err!=nil{
       	   panic(err)
	    }
		fmt.Println(index)
		srcs=append(srcs,minio.NewSourceInfo(bucketName,fmt.Sprintf("sheet_%d",i),nil))
	}
	fmt.Println(srcs)
	dst, err := minio.NewDestinationInfo("test", "222.log", nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dst)
	err = minioClient.ComposeObject(dst, srcs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Composed object successfully.")
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
	abuf := bytes.Repeat([]byte("aaaaa"), 1024*1024)//至少要5MB 才能使用分片上传(不包含最后一片)
	fmt.Println(len(abuf))
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

func TestConsolidatedFile(t *testing.T)  {
	bucketName:="test"
	var srcs =make([]minio.SourceInfo,0)
	for i:=0;i<3;i++{
		name:=fmt.Sprintf("sheet_%d",i)
		bbuf := bytes.Repeat([]byte(fmt.Sprint(i+1)), 1)//至少要5M
		index,err:=minioClient.PutObject(bucketName,name,bytes.NewReader(bbuf),int64(len(bbuf)),minio.PutObjectOptions{})
		if err!=nil{
			panic(err)
		}
		fmt.Println(index)
		srcs=append(srcs,minio.NewSourceInfo(bucketName,fmt.Sprintf("sheet_%d",i),nil))
	}
}
