package match

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

type Match struct {
	Uid       int
	Rating    int
	StartTime int64 //单位豪秒
}

type MatchPool struct {
	allUser *sync.Map
	timeout int64 //匹配超时时间  单位豪秒
	mu      sync.Mutex
	num     int //需要匹配人数
	isWork  bool
	count   int64 //匹配成功次数
}

func NewMatchPool(t int64) *MatchPool {
	pool := &MatchPool{
		allUser: new(sync.Map),
		timeout: t,
	}
	go pool.run()
	return pool
}
func (m *MatchPool) Add(u *Match) {
	m.allUser.Store(strconv.Itoa(u.Uid), u)
}

func (m *MatchPool) Remove(id int) {
	m.allUser.Delete(strconv.Itoa(id))
}

func (m *MatchPool) run() {
	for {
		select {
		case <-time.After(time.Second * 1):
			{
				m.match()
			}
		}
	}
}

func (m *MatchPool) match() {
	if m.isWork {
		return
	}
	m.mu.Lock()
	m.count++
	m.isWork = true
	defer func() {
		m.mu.Unlock()
		m.isWork = false
	}()
	fmt.Println("开始匹配时间", time.Now().UnixNano()/1e6)
	//给每个分段都添加分别的队列 （这里可以设置区间）
	var ratingMap sync.Map
	m.allUser.Range(func(_, value interface{}) bool {
		user := value.(*Match)
		if time.Now().Unix()-user.StartTime > m.timeout { //该用户匹配时间超时，剔除队列
			m.Remove(user.Uid)
		} else {
			//加入对应的分数队列
			valRating, ok := ratingMap.Load(user.Rating)
			if ok {
				val := valRating.([]Match)
				val = append(val, *user)
				//进行排序
				sort.Slice(val, func(i, j int) bool {
					return val[i].StartTime < val[j].StartTime
				})
				ratingMap.Store(user.Rating, val)
			} else {
				var userArray []Match
				userArray = append(userArray, *user)
				ratingMap.Store(user.Rating, userArray)
			}
		}
		return true
	})
	ratingMap.Range(func(rating, value interface{}) bool {
		//找出同一分数段里，等待时间最长的玩家
		continueMatch := true
		for continueMatch {
			userArray := value.([]Match)
			if len(userArray) > 0 {
				var MatchUser []Match
				maxUser := userArray[0]
				MatchUser = append(MatchUser, maxUser)
				fmt.Println("用户 UID", maxUser.Uid, "是分数", maxUser.Rating, " 上等待最久的玩家", "已经等待时间 ", time.Now().UnixNano()/1e6-maxUser.StartTime, "开始匹配时间 ", time.Now().UnixNano()/1e6)
				//先从本分数段上取数据
				for _, v := range userArray {
					if v.Uid == maxUser.Uid {
						continue
					}
					MatchUser = append(MatchUser, v)
					if len(MatchUser) >= m.num {
						break
					}
				}
				if len(MatchUser) >= m.num { //人员已经够了,不再判断
					continue
				}
				//再上下每次加1分取 如果加到50都没成功者失败

			} else {
				continueMatch = false //该分段没有数据
			}
		}
		return true
	})

}
