package bitmap_filter

import (
	"context"
	"github.com/go-redis/redis"
	"hash/fnv"
)

type BitMapFilter struct {
	conn *redis.Client
	key  string
}

func NewBitMapFileTer(conn *redis.Client, key string) *BitMapFilter {
	return &BitMapFilter{conn: conn, key: key}
}

func (b *BitMapFilter) Add(str string) error {
	key, err := hashKey(str)
	if err != nil {
		return err
	}
	return b.conn.SetBit(context.Background(), b.key, key, 1).Err()
}

func (b *BitMapFilter) Exist(str string) (bool, error) {
	key, err := hashKey(str)
	if err != nil {
		return false, err
	}
	res, err := b.conn.GetBit(context.Background(), b.key, key).Result()
	if err != nil {
		return false, err
	}
	if res != 1 {
		return false, err
	}
	return true, nil
}

func hashKey(host string) (int64, error) {
	a := fnv.New32a()
	if _, err := a.Write([]byte(host)); err != nil {
		return 0, err
	}
	return int64(a.Sum32()), nil
}
