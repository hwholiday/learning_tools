package main

import (
	"fmt"
	"testing"
)

func Test_interface(t *testing.T) {
	type Student struct {
		Name string
	}
	var b interface{} = Student{
		Name: "aaa",
	}
	fmt.Printf("%p \n", &b)
	var c = b.(Student) //复制一份给c 不影响b
	fmt.Printf("%p \n", &c)
	c.Name = "bbb"
	fmt.Println(b.(Student).Name)
	fmt.Println(c.Name)
}
