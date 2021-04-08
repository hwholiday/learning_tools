package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
)

var w sync.WaitGroup
var dirPath string
var startPage = flag.Int("s", 1, "从第几页开始爬取数据")
var endPage = flag.Int("e", 1, "爬取到第几页")

func main() {
	flag.Parse()
	d := `&q%5Bby_full_name%5D=&q%5Bby_gender%5D=any&q%5Bby_max_age%5D=&q%5Bby_min_age%5D=&q%5Bliving_in_id%5D%5B%5D=171&q%5Bliving_in_id%5D%5B%5D=171&q%5Blooking_for%5D=any&q%5Bnationality_id%5D=any`
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	dirPath = path.Join(pwd, "images")
	_, err = os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dirPath, os.ModePerm)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	if *startPage < 1 {
		fmt.Println("开始页不能小于 1")
		return
	}
	if *endPage < *startPage {
		fmt.Println("输入参数错误结束页小于开始页")
		return
	}
	var num = (*endPage-*startPage)*6 + 6
	fmt.Println(fmt.Sprintf("从 %d 页到 %d 页预计数据量 %d", *startPage, *endPage, num))
	fmt.Println("输出路径", dirPath)
	for i := *startPage; i <= *endPage; i++ {
		url := fmt.Sprintf("https://community.justlanded.com/zh/friend_finder?location=&page=%d%s", i, d)
		log.Println("爬到多少页了 : ", i)

		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
			continue
		}
		if res.StatusCode != 200 {
			continue
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		res.Body.Close()
		var img []string
		var na []string
		doc.Find(".img-fluid").Each(func(i int, s *goquery.Selection) {
			href, exist := s.Attr("src")
			if exist {
				//go downImage(href)
				img = append(img, href)
			}
		})
		doc.Find(".user-link-name").Each(func(i int, selection *goquery.Selection) {
			na = append(na, selection.Text())
		})
		if len(img) == 6 && len(na) == 6 {
			for j := 0; j < 6; j++ {
				w.Add(1)
				go downImage(img[j], na[j])
			}
			w.Wait()
		}
		if i == *endPage {
			log.Println("爬取完成")
		}
	}
	select {}
}

func downImage(url string, n1 string) {
	defer w.Done()
	n := fmt.Sprintf("%s/%s.jpg", dirPath, n1)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = ioutil.WriteFile(n, data, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
}
