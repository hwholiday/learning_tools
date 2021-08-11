package bitmap_filter

import (
	"context"
	"github.com/go-redis/redis"
	"hash/crc32"
)

type BitMapFilter struct {
	conn *redis.Client
	key  string
}

func NewBitMapFileTer(conn *redis.Client, key string) *BitMapFilter {
	return &BitMapFilter{conn: conn, key: key}
}

func (b *BitMapFilter) Add(str string) error {
	return b.conn.SetBit(context.Background(), b.key, hashKey(str), 1).Err()
}

func (b *BitMapFilter) Exist(str string) (bool, error) {
	res, err := b.conn.GetBit(context.Background(), b.key, hashKey(str)).Result()
	if err != nil {
		return false, err
	}
	if res != 1 {
		return false, err
	}
	return true, nil
}

func hashKey(host string) int64 {
	return int64(crc32.ChecksumIEEE([]byte(host)))
}
