package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//可以作为 hbase rowkey 前置补0
func TestFF(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 256; i++ {
		a := fmt.Sprintf("%02x", rand.Intn(256))
		t.Log(a)
	}
}
