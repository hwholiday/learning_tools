package service_load_balancing

import (
	"math/rand"
	"testing"
	"time"
)

var endpoints = []string{
	"127.0.0.1:80",
	"127.0.0.2:80",
	"127.0.0.3:80",
	"127.0.0.4:80",
	"127.0.0.5:80",
	"127.0.0.6:80",
	"127.0.0.7:80",
}

//从数学上得到过证明的还是经典的fisher-yates算法，主要思路为每次随机挑选一个值，放在数组末尾。然后在n-1个元素的数组中再随机挑选一个值，放在数组末尾，以此类推。
func fisher_yates(slice []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(slice); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		slice[lastIdx], slice[idx] = slice[idx], slice[lastIdx]
	}
}

func TestShuffle(t *testing.T) {
	var indexArray = make([]int, 0, len(endpoints))
	for i := 0; i < len(endpoints); i++ {
		indexArray = append(indexArray, i)
	}
	t.Log("索引数组", indexArray)
	for i := 0; i < 10; i++ {
		fisher_yates(indexArray)
		t.Log("fisher_yates 后选择的负载均衡节点信息", endpoints[indexArray[0]])

	}

}
