package main

import (
	"bytes"
	"fmt"
	"github.com/minio/minio-go"
	"io/ioutil"
	"net/url"
	"testing"
	"time"
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
func TestPutObject(t *testing.T) {
	/*_,err:=minioClient.PutObject("test","3",strings.NewReader("3"),int64(len("1")),minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}*/
	info, err := minioClient.StatObject("test", "1", minio.StatObjectOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
	fmt.Println(info.ETag)
}

func TestComposeObject(t *testing.T) {
	bucketName := "test"
	var srcs = make([]minio.SourceInfo, 0)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("sheet_%d", i)
		bbuf := bytes.Repeat([]byte(fmt.Sprint(i+1)), 1024*1025*5) //至少要5M

		index, err := minioClient.PutObject(bucketName, name, bytes.NewReader(bbuf), int64(len(bbuf)), minio.PutObjectOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Println(index)
		srcs = append(srcs, minio.NewSourceInfo(bucketName, fmt.Sprintf("sheet_%d", i), nil))
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
	abuf := bytes.Repeat([]byte("aaaaa"), 1024*1024) //至少要5MB 才能使用分片上传(不包含最后一片)
	fmt.Println(len(abuf))
	bbuf := bytes.Repeat([]byte("b"), 1024*1024)
	daya := minio.Core{Client: minioClient}
	bucket := "test"
	name := "3.log"
	v1, err := daya.NewMultipartUpload(bucket, name, minio.PutObjectOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(v1)
	var completeParts []minio.CompletePart
	part, err := daya.PutObjectPart(bucket, name, v1, 1, bytes.NewReader(abuf), int64(len(abuf)), "", "", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("1", part)
	completeParts = append(completeParts, minio.CompletePart{PartNumber: part.PartNumber, ETag: part.ETag})

	part, err = daya.PutObjectPart(bucket, name, v1, 2, bytes.NewReader(bbuf), int64(len(bbuf)), "", "", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("2", part.ETag)
	fmt.Println("2", part.PartNumber)
	fmt.Println("2", part.LastModified)
	fmt.Println("2", part.Size)
	completeParts = append(completeParts, minio.CompletePart{PartNumber: part.PartNumber, ETag: part.ETag})
	_, err = daya.CompleteMultipartUpload(bucket, name, v1, completeParts)
	if err != nil {
		panic(err)
	}
}

func TestMinio(t *testing.T) {
	var (
		bucket = "bat-app"
	)
	minioClient, err := minio.New("192.168.2.28:9090", "672I9BRC71NVDCMASYQQ", "yl1AGr-ViO6queLv93EFlRv-iZ1icSvUVE8j6pi9", false)
	if err != nil {
		panic(err)
	}
	if err = minioClient.MakeBucket(bucket, ""); err != nil {
		panic(err)
	}
	fmt.Println("创建桶成功:" + bucket)
	//设置桶的权限
	policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:ListBucketMultipartUploads","s3:GetBucketLocation","s3:ListBucket"],"Resource":["arn:aws:s3:::abc-test-1"]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:ListMultipartUploadParts","s3:PutObject","s3:AbortMultipartUpload","s3:DeleteObject","s3:GetObject"],"Resource":["arn:aws:s3:::abc-test-1/*"]}]}`
	err = minioClient.SetBucketPolicy(bucket, policy)
	if err != nil {
		fmt.Println(err)
		return
	}

	//判断文件是否存在
	fmt.Println(minioClient.StatObject(bucket, "conf.txt", minio.StatObjectOptions{}))

	//生成上传链接
	expiry := time.Second * 24 * 60 * 60 // 1 day.
	presignedURL, err := minioClient.PresignedPutObject(bucket, "PayDemo.java", expiry)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)
	reqParams := make(url.Values)
	reqParams.Set("Content-Type", "image/jpeg")
	presignedURL, err = minioClient.Presign("PUT", bucket, "20190306121434.jpg", time.Second*24*60*60, reqParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully generated presigned URL \n", presignedURL)
}
