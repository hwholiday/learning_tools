package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	Name string
	time.Time
	Err error
}

func (m MyError) Error() string {
	return fmt.Sprintf("%v %v %v", m.Name, m.Time, m.Err)
}
func (e *MyError) Unwrap() error {
	return e.Err
}

func TestErr() error {
	return MyError{
		Name: "AAAA",
		Time: time.Now(),
	}
}


func main() {
	err := Test()
	fmt.Println("产生的错误",err)
	var testErr MyError
	fmt.Println("解析错误内容",errors.As(err, &testErr)) //查询err里面是否有自定义的MyError错误,并解除其中数据
	fmt.Println(testErr)
	fmt.Println("判断是否有该错误",errors.Is(err, testErr)) //是否包含该错误
	fmt.Println("去掉最上的错误",errors.Unwrap(err))

}

func Test() error {
	err := TestErr()
	err = fmt.Errorf("加入第一个错误%w", err)
	return err
}
