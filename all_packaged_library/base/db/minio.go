package db

import (
	"file_storage/base/config"
	"file_storage/base/tool"
	"github.com/minio/minio-go"
)

func initMinio() {
	var secure bool
	if config.GetMinioConfig().GetPath() == "s3.amazonaws.com" {
		secure = true
	}
	if minioClient, err = minio.New(config.GetMinioConfig().GetPath(), config.GetMinioConfig().GetAccessKeyId(), config.GetMinioConfig().GetSecretAccessKey(), secure); err != nil {
		panic(err)
	}
	tool.GetLogger().Debug("minio success : " + config.GetMinioConfig().GetPath())
	/*for i:=1;i<=100;i++{
		var bucketName bytes.Buffer
		bucketName.WriteString("storage")
		bucketName.WriteString("-")
		bucketName.WriteString(strconv.Itoa(int(i)))
		fmt.Println(minioClient.MakeBucket(bucketName.String(),""))
	}*/
}
