package main

import (
	"bytes"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/minio/minio-go"
	"os"
	"time"
)

var (
	minioCore   *minio.Core
	redisClient *redis.Client
	err         error
)

func InitMinio() {
	minioCore, err = minio.NewCore("127.0.0.1:9000", "AKIAINCHU2DIYAQ66TPA", "ZaYtz0d61fJXQ7djyXBX4yZ5ob8Kj/WNXtw6PJob", false)
	CheckErr("minio.New", err)
}

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	pong, err := redisClient.Ping().Result()
	CheckErr(pong, err)
}

//分片上传文件
//分片文件小于5M 再redis缓存到5M在上传
func main() {
	InitMinio()
	InitRedis()
	var (
		bucket        = "test"
		name          = "3.log"
		uploadId      string
		completeParts []minio.CompletePart
		maxSize       = 1024 * 1024 * 5
	)

	uploadId, err = minioCore.NewMultipartUpload(bucket, name, minio.PutObjectOptions{})
	CheckErr("NewMultipartUpload", err)
	fmt.Println("uploadId", uploadId)
	var j = 1
	for i := 1; i < 20; i++ {
		//先判断文件大小是不是满足5M
		time.Sleep(time.Second * 3)
		a := bytes.Repeat([]byte(fmt.Sprint(1)), 1024*1024)
		redisA, err := redisClient.Get(uploadId).Bytes()
		if err != nil {
			if err != redis.Nil {
				fmt.Println("redisClient.Get", uploadId, err)
				panic(err)
			}
		}
		var data bytes.Buffer
		data.Write(redisA)
		data.Write(a)
		fmt.Println("index", i, "  文件大小", len(data.Bytes())/1024/1024)
		if len(data.Bytes()) < maxSize { //小于继续缓存
			if i >= 19 { //最后一片
				//最后一片直接上传
				fmt.Println("最后一片直接上传")
			} else {
				fmt.Println(uploadId, "内容太小继续缓存")
				redisClient.Set(uploadId, data.Bytes(), time.Minute*30)
				continue
			}

		}
		part, err := minioCore.PutObjectPart(bucket, name, uploadId, j, bytes.NewReader(data.Bytes()), int64(len(data.Bytes())), "", "", nil)
		CheckErr("PutObjectPart", err)
		completeParts = append(completeParts, minio.CompletePart{PartNumber: part.PartNumber, ETag: part.ETag})
		fmt.Println("uploadId :", uploadId, "part :", j, part)
		//该缓存已经上传完毕删除该内容
		fmt.Println("uploadId :", uploadId, redisClient.Del(uploadId).String())
		j++

	}
	_, err = minioCore.CompleteMultipartUpload(bucket, name, uploadId, completeParts)
	if err != nil {
		panic(err)
	}
	fmt.Println("上传文件成功", uploadId)
}

func CheckErr(msg string, err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("%s :%v", msg, err.Error()))
		os.Exit(1)
	}
}

//uploadId 27ce1dec-f742-433e-a284-8eafeb22476b
//index 1   文件大小 1
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 2   文件大小 2
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 3   文件大小 3
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 4   文件大小 4
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 5   文件大小 5
//uploadId : 27ce1dec-f742-433e-a284-8eafeb22476b part : 1 {1 0001-01-01 00:00:00 +0000 UTC 7f0883388269a3eeb5b116b39ad81ee9-1 5242880}
//uploadId : 27ce1dec-f742-433e-a284-8eafeb22476b del 27ce1dec-f742-433e-a284-8eafeb22476b: 1
//index 6   文件大小 1
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 7   文件大小 2
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 8   文件大小 3
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 9   文件大小 4
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 10   文件大小 5
//uploadId : 27ce1dec-f742-433e-a284-8eafeb22476b part : 2 {2 0001-01-01 00:00:00 +0000 UTC 4de9698076020ddd32de7230316d8ce5-1 5242880}
//uploadId : 27ce1dec-f742-433e-a284-8eafeb22476b del 27ce1dec-f742-433e-a284-8eafeb22476b: 1
//index 11   文件大小 1
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 12   文件大小 2
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 13   文件大小 3
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 14   文件大小 4
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 15   文件大小 5
//uploadId : 27ce1dec-f742-433e-a284-8eafeb22476b part : 3 {3 0001-01-01 00:00:00 +0000 UTC 3717c6e04fe51d8028344be6059acef7-1 5242880}
//uploadId : 27ce1dec-f742-433e-a284-8eafeb22476b del 27ce1dec-f742-433e-a284-8eafeb22476b: 1
//index 16   文件大小 1
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 17   文件大小 2
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 18   文件大小 3
//27ce1dec-f742-433e-a284-8eafeb22476b 内容太小继续缓存
//index 19   文件大小 4
//最后一片直接上传
//uploadId : 27ce1dec-f742-433e-a284-8eafeb22476b part : 4 {4 0001-01-01 00:00:00 +0000 UTC e4c04a2710241b58b66fe51677acf3af-1 4194304}
//uploadId : 27ce1dec-f742-433e-a284-8eafeb22476b del 27ce1dec-f742-433e-a284-8eafeb22476b: 1
//上传文件成功 27ce1dec-f742-433e-a284-8eafeb22476b
