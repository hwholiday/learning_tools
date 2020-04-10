package match

import (
	"fmt"
	"math"
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

func NewMatchPool(t int64, num int) *MatchPool {
	pool := &MatchPool{
		allUser: new(sync.Map),
		timeout: t,
		num:     num,
	}
	return pool
}
func (m *MatchPool) Add(u *Match) {
	m.allUser.Store(strconv.Itoa(u.Uid), u)
}

func (m *MatchPool) Remove(id int) {
	m.allUser.Delete(strconv.Itoa(id))
}

func (m *MatchPool) Run() {
	m.match()
	//for {
	//	select {
	//	case <-time.After(time.Second * 1):
	//		{
	//			m.match()
	//		}
	//	}
	//}
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
	fmt.Println("开始匹配时间", time.Now().UnixNano()/1e6, "次数", m.count)
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
	ratingMap.Range(func(key, value interface{}) bool {
		m.matchUser(ratingMap, key, value)
		return true
	})

}

func (m *MatchPool) matchUser(ratingMap sync.Map, key, value interface{}) {
	//找出同一分数段里，等待时间最长的玩家
	/*	continueMatch := true
		for continueMatch {*/
	userArray := value.([]Match)
	if len(userArray) <= 0 {
		return
	}
	var MatchUser []Match
	maxUser := userArray[0] //等待时间最长的用户
	MatchUser = append(MatchUser, maxUser)
	m.allUser.Delete(maxUser.Uid)
	userArray = append(userArray[:0], userArray[0+1:]...)
	ratingMap.Store(key, userArray)
	m.allUser.Delete(maxUser.Uid)
	fmt.Println("用户 UID", maxUser.Uid, "是分数", maxUser.Rating, " 上等待最久的玩家", "已经等待时间 ", time.Now().UnixNano()/1e6-maxUser.StartTime, "开始匹配时间 ", time.Now().UnixNano()/1e6)
	waitSecond := time.Now().Unix() - maxUser.StartTime/1000
	step := getRatingStep(waitSecond)
	min := maxUser.Rating - step
	if min < 0 {
		min = 0
	}
	max := maxUser.Rating + step
	fmt.Println("用户 UID ", maxUser.Uid, "本次搜索 rating 范围下限 ", min, "rating 范围上限 ", max)
	//TODO 再上下每次加1分取 如果加到都没成功者失败
	middle := maxUser.Rating //设置 rating 区间中位数
	for searchRankUp, searchRankDown := middle, middle; searchRankUp <= max && searchRankDown >= min; searchRankUp, searchRankDown = searchRankUp+1, searchRankDown-1 {
		if searchRankDown != searchRankUp && searchRankDown > 0 { //目前只选择比我评分低的人，体验会好一些
			rank, ok := ratingMap.Load(searchRankDown)
			if !ok {
				continue
			}
			rankArray := rank.([]Match)
			if len(rankArray) <= 0 {
				continue
			}
			for index, v := range rankArray {
				if v.Uid != maxUser.Uid {
					if len(MatchUser) < m.num {
						MatchUser = append(MatchUser, v)
						fmt.Println("用户 UID ", maxUser.Uid, "在  rating", searchRankDown, " 找到匹配用户  ", v.Uid)
						//移除该数据
						rankArray = append(rankArray[:index], rankArray[index+1:]...)
						ratingMap.Store(searchRankDown, rankArray)
						m.allUser.Delete(v.Uid)
					} else {
						break
					}
				}
			}
			if len(MatchUser) >= m.num { //匹配人员成功
				fmt.Println("匹配到人 ", MatchUser)
			} else { //本地匹配失败
				fmt.Println("这个分段匹配失败 ", key)
				m.allUser.Store(maxUser.Uid, maxUser)
				/*	continueMatch = false*/
			}
		}
		/*	}*/
	}
}

//waitSecond 该用户等待了多少秒
func getRatingStep(waitSecond int64) int {
	var (
		step             = 1.3
		baseStep float64 = 3
		maxStep  float64 = 100
	)
	u := math.Pow(float64(waitSecond), step)
	u = u + baseStep
	u = math.Round(u)
	u = math.Min(u, maxStep) //等待时间越长，rating 区间越大
	return int(u)
}
