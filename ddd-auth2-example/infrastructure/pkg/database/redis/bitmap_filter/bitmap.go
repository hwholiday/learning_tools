package bitmap_filter

import (
	"errors"
	"github.com/go-redis/redis"
)

type BitMapFilter struct {
	conn redis.Cmdable
	key  string
}

func NewBitMapFileTer(conn redis.Cmdable, key string) *BitMapFilter {
	return &BitMapFilter{conn: conn, key: key}
}

func (b *BitMapFilter) Add(str string) error {
	return b.conn.HSet(b.key, str, 1).Err()
}

func (b *BitMapFilter) Exist(str string) (bool, error) {
	res, err := b.conn.HGet(b.key, str).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	if res != "1" {
		return false, nil
	}
	return true, nil
}
