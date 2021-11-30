package repository

import (
	"context"
	"encoding/json"
	"github.com/hwholiday/learning_tools/ddd-auth2-example/infrastructure/pkg/database/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	mgo *mongo.Client
	rds *redis.Client
}

type Repository struct {
	Merchant  *Merchant
	AuthCode  *AuthCode
	AuthToken *AuthToken
}

func NewRepository(mgo *mongo.Client, r *redis.Client) *Repository {
	var a = repository{mgo, r}
	return &Repository{
		&Merchant{a},
		&AuthCode{a},
		&AuthToken{a},
	}
}

func (r *repository) Close() {
	if r.mgo != nil {
		_ = r.mgo.Disconnect(context.Background())
	}
	if r.rds != nil {
		_ = r.rds.Close()
	}
}

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
