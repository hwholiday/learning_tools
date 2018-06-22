package main

import (
	"github.com/minio/minio-go"
	"log"
	"net/http"
	"fmt"
	"time"
	"sync"
)

const (
	MINIO_ACCESSKEY = "1J0FOOJFC899D3FG3YEM"
	MINIO_SECRETKEY = "ZzuccGQMFBFTFudS9QfUFBoqIk1ZnujBArHKXTYB"
	MINIO_ENDPOINT  = "127.0.0.1:9000"
	MINIO_USERSSL   = false
	bucket          = "howie"
	FILE_TIMEOUT    = time.Hour * 24
)

var MinioClient *minio.Client
var FileMap = sync.Map{}

func init() {
	var err error
	MinioClient, err = minio.New(MINIO_ENDPOINT, MINIO_ACCESSKEY, MINIO_SECRETKEY, MINIO_USERSSL)
	if err != nil {
		log.Panic(err)
	}
	err = MinioClient.MakeBucket(bucket, "")
	if err != nil {
		exists, err := MinioClient.BucketExists(bucket)
		if exists && err == nil {
			//log.Println(bucket + "已经存在")
			return
		} else {
			log.Panic(err)
		}
	}
	log.Println("成功创建:" + bucket)
}
func main() {
	http.HandleFunc("/v1/up", upFile)
	http.HandleFunc("/v1/find", findFile)
	http.HandleFunc("/v1/remove", removeFile)
	fmt.Println("服务器启动:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
//上传文件
func upFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("请使用POST方法"))
		return
	}
	fileinfo, hander, err := r.FormFile("up_file")
	defer fileinfo.Close()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	n, err := MinioClient.PutObject(bucket, hander.Filename, fileinfo, hander.Size, minio.PutObjectOptions{ContentType: hander.Header.Get("Content-Type")})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if hander.Size != n {
		w.Write([]byte("上传失败"))
		return
	}
	id := fmt.Sprint(time.Now().UnixNano())
	FileMap.Store(id, hander.Filename)
	w.Write([]byte("上传成功,资源ID为:" + id, ))
	return
}

//查找文件
func findFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("请使用GET方法"))
		return
	}
	r.ParseForm()
	id := r.Form.Get("id")
	if id == "" {
		w.Write([]byte("请传入资源ID"))
		return
	}
	v, ok := FileMap.Load(id)
	if !ok {
		w.Write([]byte("没有找到该资源"))
		return
	}
	info, err := MinioClient.PresignedGetObject(bucket, fmt.Sprint(v), FILE_TIMEOUT, nil)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(info.String()))
	return
}

//删除文件
func removeFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Write([]byte("请使用DELETE方法"))
		return
	}
	r.ParseForm()
	id := r.Form.Get("id")
	if id == "" {
		w.Write([]byte("请传入资源ID"))
		return
	}
	v, ok := FileMap.Load(id)
	if !ok {
		w.Write([]byte("没有找到该资源"))
		return
	}
	if err := MinioClient.RemoveIncompleteUpload(bucket, fmt.Sprint(v)); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	FileMap.Delete(id)
	w.Write([]byte("删除成功"))
	return
}

