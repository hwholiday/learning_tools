package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/robfig/cron"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

var filePath = flag.String("p", "./test.log", "nginx log path")
var urlIndex = flag.Int("i", 5, "url index")
var cronTime = flag.String("t", "*/5 * * * * ?", "cron Time")
var statisticsUrl map[string]int
var mu sync.Mutex
var listArray []list

type list struct {
	key string
	num int
}

func main() {
	flag.Parse()
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	statisticsUrl = make(map[string]int)
	t, err := tail.TailFile(*filePath, config)
	if err != nil {
		fmt.Println("TailFile", err)
		return
	}
	go func() {
		c := cron.New()
		if err := c.AddFunc(*cronTime, func() {
			f, err := os.Create(fmt.Sprintf("./nginx_statistics_%v.log", time.Now().Format("2006-01-02_15:04:05")))
			if err != nil {
				fmt.Println("Create", err)
				return
			}
			defer f.Close()
			w := bufio.NewWriter(f)
			mu.Lock()
			for k, v := range statisticsUrl {
				listArray = append(listArray, list{
					key: k,
					num: v,
				})
			}
			sort.Slice(listArray, func(i, j int) bool {
				return listArray[i].num > listArray[j].num
			})
			for _, v := range listArray {
				_, _ = fmt.Fprintln(w, fmt.Sprintf("%d     %s", v.num, v.key))
			}
			_ = w.Flush()
			statisticsUrl = make(map[string]int)
			listArray = []list{}
			mu.Unlock()
		}); err != nil {
			fmt.Println("AddFunc", err)
			return
		}
		c.Start()
	}()
	for {
		select {
		case info := <-t.Lines:
			i := strings.Split(info.Text, " ")
			for index, v := range i {
				if index == *urlIndex {
					mu.Lock()
					val, ok := statisticsUrl[v]
					if ok {
						val = val + 1
						statisticsUrl[v] = val
					} else {
						statisticsUrl[v] = 1
					}
					mu.Unlock()
				}
			}
		}
	}
}
