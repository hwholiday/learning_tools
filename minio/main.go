package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"os"
	"time"
)

func main() {
	var (
		bucket = "abc-test"
	)
	minioClient, err := minio.New("192.168.2.28:9090", "Z5OID6N9JVAY0CKL40FS", "YtP3xC9CWfSecFJzzm5gP3GwpjISZwaWotzeL6Se", false)
	CheckErr("minio.New", err)
	/*if err = minioClient.MakeBucket(bucket, ""); err != nil {
		CheckErr("创建储存桶", err)
	}
	fmt.Println("创建桶成功:" + bucket)
	//设置桶的权限
	policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:ListBucketMultipartUploads","s3:GetBucketLocation","s3:ListBucket"],"Resource":["arn:aws:s3:::abc-test-1"]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:ListMultipartUploadParts","s3:PutObject","s3:AbortMultipartUpload","s3:DeleteObject","s3:GetObject"],"Resource":["arn:aws:s3:::abc-test-1/*"]}]}`
	err = minioClient.SetBucketPolicy(bucket, policy)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	//判断文件是否存在
	fmt.Println(minioClient.StatObject(bucket, "conf.txt1", minio.StatObjectOptions{}))

	//生成上传链接
	expiry := time.Second * 24 * 60 * 60 // 1 day.
	presignedURL, err := minioClient.PresignedPutObject(bucket, "PayDemo.java", expiry)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)
}

func CheckErr(msg string, err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("%s :%v", msg, err.Error()))
		os.Exit(1)
	}
}
