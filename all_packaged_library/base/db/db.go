package db

import (
	"github.com/go-redis/redis"
	"github.com/minio/minio-go"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var (
	err         error
	redisDb     *redis.Client
	minioClient *minio.Client
	mgo         *mongo.Client
	m           sync.Mutex
)

func Init() {
	m.Lock()
	defer m.Unlock()
	initRedis()
	initMongoDb()
	initMinio()
}

func GetRedisDb() *redis.Client {
	return redisDb
}

func CloseRedisDb() {
	closeRedis()
}

func GetMgoDb() *mongo.Client {
	return mgo
}

func GetMinioClient() *minio.Client {
	return minioClient
}
