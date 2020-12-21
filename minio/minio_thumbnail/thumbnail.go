package main

import (
	"bytes"
	"fmt"
	"github.com/minio/minio-go"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"strings"
)

var minioClient *minio.Client

func StartMinio() {
	var err error
	// 初使化 minio client对象。
	minioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		panic(err)
		return
	}
	doneCh := make(chan struct{})
	defer close(doneCh)
	fmt.Println("server start success")
	for notificationInfo := range minioClient.ListenBucketNotification("images", "", ".jpg", []string{
		"s3:ObjectCreated:*",
	}, doneCh) {
		if notificationInfo.Err != nil {
			log.Fatalln(notificationInfo.Err)
		}
		for _, v := range notificationInfo.Records {
			bucket := v.S3.Bucket.Name
			fileName := v.S3.Object.Key
			go GetFileAndThumbnail(bucket, fileName)
		}
	}
}

func GetFileAndThumbnail(bucket, fileName string) {
	if strings.Contains(fileName, "thumbnail") {
		return
	}
	//log.Println("[GetFileAndThumbnail] Start", "bucket", bucket, "fileName", fileName)
	//判断是不是图片
	n := strings.Split(fileName, ".")
	if !CheckIsImage(n[1]) {
		return
	}
	thumbnailName := fmt.Sprintf("%s_%s.%s", n[0], "thumbnail", "jpg")
	if CheckFileEx(bucket, thumbnailName) {
		return
	}
	if len(strings.Split(fileName, ".")) != 2 {
		log.Println("[GetFileAndThumbnail] fileName != 2", fileName, "bucket", bucket)
		return
	}
	reader, err := minioClient.GetObject(bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Println("[GetFileAndThumbnail] bucket", bucket, "fileName", fileName, "GetObject", err)
		return
	}
	defer reader.Close()
	readerByte, err := jpeg.Decode(reader)
	if err != nil {
		log.Println("[GetFileAndThumbnail] bucket", bucket, "fileName", fileName, " ReadAll", err)
		return
	}
	m := resize.Thumbnail(maxWidth, maxHeight, readerByte, resize.Lanczos3)
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, m, nil)
	if err != nil {
		log.Println("[GetFileAndThumbnail] bucket", bucket, "fileName", fileName, " Encode", err)
		return
	}
	if CheckFileEx(bucket, thumbnailName) {
		return
	}
	_, err = minioClient.PutObject(bucket, thumbnailName, bytes.NewReader(buf.Bytes()), int64(len(buf.Bytes())), minio.PutObjectOptions{
		ContentType: "image/jpg",
	})
	if err != nil {
		log.Println("[GetFileAndThumbnail] bucket", bucket, "fileName", "fileName", " PutObject", err)
		return
	}
	//log.Println("[GetFileAndThumbnail] End bucket", bucket, "fileName", fileName, "thumbnail", thumbnailName)
}

func CheckFileEx(bucket, fileName string) bool {
	objInfo, err := minioClient.StatObject(bucket, fileName, minio.StatObjectOptions{})
	if err != nil {
		if err.Error() != "The specified key does not exist." {
			fmt.Println("[GetFileAndThumbnail] CheckFileEx  bucket", bucket, "fileName", fileName, "err", err)
		}
		return false
	}
	return objInfo.Key != ""
}

func CheckIsImage(s string) bool {
	s = strings.ToUpper(s)
	for _, v := range checkImage {
		if v == s {
			return true
		}
	}
	return false
}
