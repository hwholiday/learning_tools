package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"os"
	"reflect"
	"testing"
	"time"
)

var client *elastic.Client

func TestMain(m *testing.M) {
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://172.13.3.160:9200/"))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping("http://172.13.3.160:9200/").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	esversion, err := client.ElasticsearchVersion("http://172.13.3.160:9200/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
	os.Exit(m.Run())
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

var indexName = "hwholiday"

func getId() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
}

func TestCreate(t *testing.T) {
	res, err := client.Index().Index(indexName).Id("1").BodyJson(User{
		FirstName: "a",
		LastName:  "b",
		Age:       10,
	}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	res, err = client.Index().Index(indexName).Id("2").BodyJson(User{
		FirstName: "a",
		LastName:  "c",
		Age:       20,
	}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	res, err = client.Index().Index(indexName).Id("3").BodyJson(User{
		FirstName: "abcd e",
		LastName:  "f",
		Age:       30,
	}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestUpdate(t *testing.T) {
	res, err := client.Update().Index(indexName).Id("1").Doc(map[string]interface{}{
		"age":       50,
		"last_name": "c",
	}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGet(t *testing.T) {
	res, err := client.Get().Index(indexName).Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	var data User
	err = json.Unmarshal(res.Source, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func TestDel(t *testing.T) {
	res, err := client.Delete().Index(indexName).Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestInsertArray(t *testing.T) {
	res, err := client.Search(indexName).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range res.Each(reflect.TypeOf(User{})) {
		t := v.(User)
		fmt.Println(t)
	}
}

func TestQuery(t *testing.T) {
	//字段相等
	q := elastic.NewQueryStringQuery("last_name:c")
	printQuery(q)
	//条件查询
	bQ := elastic.NewBoolQuery()
	bQ.Must(elastic.NewMatchQuery("last_name", "c"))
	bQ.Filter(elastic.NewRangeQuery("age").Gt(10))
	printQuery(bQ)

	//短语搜索 搜索 first_name字段中有 a
	mQ := elastic.NewMatchPhraseQuery("first_name", "e")
	printQuery(mQ)
}

func TestList(t *testing.T) {
	var (
		size = 10
		page = 1
	)
	if size < 0 || page < 1 {
		return
	}
	res, err := client.Search(indexName).
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range res.Each(reflect.TypeOf(User{})) {
		t := v.(User)
		fmt.Println(t)
	}
}

func printQuery(q elastic.Query) {
	res, err := client.Search(indexName).Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range res.Each(reflect.TypeOf(User{})) {
		t := v.(User)
		fmt.Println(t)
	}
}
