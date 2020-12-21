package hbase

import (
	"context"
	"fmt"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
	"time"
)

var client gohbase.Client

//每一秒产生一个
func GetRowKey(uid uint32) string {
	return fmt.Sprintf("%s%d", strconv.Itoa(int(uid)), math.MaxInt64-time.Now().Unix())
}
func reverse(str string) string {
	var result string
	length := len(str)
	for i := 0; i < length; i++ {
		result += fmt.Sprintf("%c", str[length-i-1])
	}
	return result
}

func TestMain(m *testing.M) {
	client = gohbase.NewClient("172.13.3.160")
	os.Exit(m.Run())
}

func TestCreateTable(t *testing.T) {
	val := map[string]map[string]string{
		"info": {
			"name": fmt.Sprintf("test%d", 2),
			"age":  fmt.Sprintf("%d", 2),
		},
	}
	hrpc.NewCreateTable(context.Background(), []byte("log_table"), val)
}

func TestCreate(t *testing.T) {
	for i := 10; i <= 60; i++ {
		val := map[string]map[string][]byte{
			"info": {
				"name": []byte(fmt.Sprintf("test%d", i)),
				"age":  []byte(fmt.Sprintf("%d", i)),
			},
		}
		rowKey := GetRowKey(uint32(i))
		t.Log("rowKey", rowKey)
		p, err := hrpc.NewPutStr(context.Background(), "user", rowKey, val)
		if err != nil {
			t.Error(err)
			return
		}
		rsp, err := client.Put(p)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(rsp.String())
	}
}

func TestGet(t *testing.T) {
	q, err := hrpc.NewGetStr(context.Background(), "user", "543219223372035248499836")
	rsp, err := client.Get(q)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rsp.String())
	for _, v := range rsp.Cells {
		val := v
		t.Log(string(val.Qualifier), string(val.Value))
	}
}

func TestDel(t *testing.T) {
	d, err := hrpc.NewDelStr(context.Background(), "user", "543219223372035248499836", nil)
	if err != nil {
		t.Error(err)
	}
	rsp, err := client.Delete(d)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rsp.String())
}

func TestGetFamily(t *testing.T) {
	//109223372035248490685                                column=info:age, timestamp=1606285123443, value=10
	//109223372035248490685                                column=info:name, timestamp=1606285123443, value=test10
	family := map[string][]string{"info": []string{"name"}}
	q, err := hrpc.NewGetStr(context.Background(), "user", "109223372035248490685", hrpc.Families(family))
	rsp, err := client.Get(q)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rsp.String())
	for _, v := range rsp.Cells {
		val := v
		t.Log(string(val.Qualifier), string(val.Value))
	}
}

func TestScanStr(t *testing.T) {
	//rowKey 前缀匹配
	pFilter := filter.NewPrefixFilter([]byte("17"))
	q, err := hrpc.NewScanStr(context.Background(), "user", hrpc.Filters(pFilter))
	if err != nil {
		t.Error(err)
		return
	}
	scanRsp := client.Scan(q)
	for {
		res, err := scanRsp.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range res.Cells {
			val := v
			t.Log(string(val.Qualifier), string(val.Value))
		}
	}
}

func TestCreate2(t *testing.T) {
	for i := 1; i <= 50; i++ {
		val := map[string]map[string][]byte{
			"info": {
				"name": []byte(fmt.Sprintf("hero%d", i)),
				"age":  []byte(fmt.Sprintf("%d", i)),
			},
		}
		rowKey := GetRowKey(uint32(123))
		t.Log("rowKey", rowKey)
		p, err := hrpc.NewPutStr(context.Background(), "user", rowKey, val)
		if err != nil {
			t.Error(err)
			return
		}
		rsp, err := client.Put(p)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(rsp.String())
		time.Sleep(time.Second)
	}
}

//分页查询
func TestScanRangeStr(t *testing.T) {
	f := filter.NewPageFilter(11)
	q, err := hrpc.NewScanRangeStr(context.Background(), "user", "1239223372035248488854", "", hrpc.Filters(f))
	if err != nil {
		t.Error(err)
		return
	}
	scanRsp := client.Scan(q)
	for {
		res, err := scanRsp.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range res.Cells {
			val := v
			t.Log(string(val.Qualifier), string(val.Value), string(val.Row))
		}
	}
}

func TestScanRangeStr2(t *testing.T) {
	// 获取这个区间的数据
	// 1606286961323  2020-11-25 14:49:21
	// 1606286952113  2020-11-25 14:49:12
	startKey := fmt.Sprintf("%d%d", 123, math.MaxInt64-1606286961)
	endKey := fmt.Sprintf("%d%d", 123, math.MaxInt64-1606286952)
	q, err := hrpc.NewScanRangeStr(context.Background(), "user", startKey, endKey)
	if err != nil {
		t.Error(err)
		return
	}
	scanRsp := client.Scan(q)
	for {
		res, err := scanRsp.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range res.Cells {
			val := v
			t.Log(string(val.Qualifier), string(val.Value), string(val.Row))
		}
	}
}
