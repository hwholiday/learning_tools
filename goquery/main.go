package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Account struct {
	Name      string
	EOS       string
	NETWeight string
	CPUWeight string
}

func main() {
	res, err := http.Get("https://eosmonitor.io/accounts?page=1")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data []string
	doc.Find(".common-lsit-data_table").Each(func(i int, s *goquery.Selection) {
		tdInfo := strings.Split(s.Find("td").Text(), "\n")
		for _, v := range tdInfo {
			if len(v) != 0 && len(v) != 12 {
				data = append(data, strings.Replace(v, " ", "", -1))
			}
		}
	})
	var accounts []Account
	var account Account
	for i, v := range data {
		switch (i + 1) % 4 {
		case 1: //Name
			account.Name = v
			break
		case 2: //EOS
			account.EOS = v
			break
		case 3: //NETWeight
			account.NETWeight = v
			break
		case 0: //CPUWeight
			account.CPUWeight = v
			accounts = append(accounts, account)
			break

		}
	}
	fmt.Println(accounts)
	fmt.Println("结束")
}
