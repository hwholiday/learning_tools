package main

import (
	"fmt"
	"reflect"
)

type Howie struct {
	Name string
	Sex  int
}

func (h Howie) Do()string {
	fmt.Println("调用该方法")
	return "调用该方法2"
}

func main() {
	var i = 1
	var ii = "11"
	var iii = []string{"1", "2", "3"}
	var h Howie
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(ii))
	fmt.Println(reflect.TypeOf(iii))
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.TypeOf(h).Name())
	hh := reflect.TypeOf(h)
	for i := 0; i < hh.NumField(); i++ {
		fmt.Println(hh.Field(i).Name)
		fmt.Println(hh.Field(i).Type)
	}
	fmt.Println("hh.Kind()",hh.Kind())
	fmt.Println("--------------------")
	fmt.Println(reflect.ValueOf(i))
	fmt.Println(reflect.ValueOf(ii))
	fmt.Println(reflect.ValueOf(iii))
	fmt.Println(reflect.ValueOf(h))
	var hhh = Howie{Name: "1111", Sex: 2}
	var rt = reflect.TypeOf(hhh)
	var rv = reflect.ValueOf(hhh)
	for i := 0; i < rt.NumField(); i++ {
		fmt.Println(hh.Field(i).Name)
		fmt.Println(rv.Field(i).Type())
		fmt.Println(rv.Field(i).Interface())
	}
 	//reflect ValueOf 赋值
	fv := reflect.ValueOf(i)
	fe := reflect.ValueOf(&i).Elem()
	fmt.Println(fv)
	fmt.Println(fe)
	fmt.Println(fe.CanSet())
	fe.SetInt(123)
	fmt.Println(i)
	fmt.Println(reflect.ValueOf(i).Interface())
	//reflect ValueOf 调用方法
	data := reflect.ValueOf(hhh).MethodByName("Do").Call([]reflect.Value{})
	fmt.Println(data[0].String())
}
