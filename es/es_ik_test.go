package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
	"testing"
)

type Data struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var indexIkName = "ik_test"

func TestIkIndex(t *testing.T) {
	data := `{
   "settings":{
		"number_of_shards":5,
		"number_of_replicas":1
	},
  "mappings": {
      "properties": {
       "id":{"type":"long"},
       "title": {
                "type": "text",
                "analyzer": "ik_smart",
            },
       "content": {
                "type": "text",
                "analyzer": "ik_max_word",
            }
      }
  }
}`
	res, err := client.CreateIndex(indexIkName).BodyJson(data).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
func TestIkCreate(t *testing.T) {
	res, err := client.Index().Index(indexIkName).Id(getId()).BodyJson(Data{
		Id:      1,
		Title:   "软件工程师",
		Content: "我们是软件工程师，我们是中国人",
	}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	res, err = client.Index().Index(indexIkName).Id(getId()).BodyJson(Data{
		Id:      1,
		Title:   "中国",
		Content: "我们是建筑工程师",
	}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestQueryIk(t *testing.T) {
	q := elastic.NewMatchQuery("content", "工程师")
	res, err := client.Search(indexIkName).Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range res.Each(reflect.TypeOf(Data{})) {
		t := v.(Data)
		fmt.Println(t)
	}
}
