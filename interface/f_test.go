package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Srv interface {
	Add(a int, b int) (int, error)
	Del(a int) error
}

type srv struct {
}

func (s *srv) Add(a int, b int) (int, error) {
	return a + b, nil
}
func (s *srv) Del(a int) error {
	return nil
}

func Test_interface(t *testing.T) {
	var handler Srv
	handler = &srv{}
	typ := reflect.TypeOf(handler)
	hdlr := reflect.ValueOf(handler)
	name := reflect.Indirect(hdlr).Type().Name()
	fmt.Println(name)
	for m := 0; m < typ.NumMethod(); m++ {
		fmt.Println("m", m, typ.Method(m).Type)
		fmt.Println("m", m, typ.Method(m).Name)
		fmt.Println("m", m, typ.Method(m).Func)
		fmt.Println("m", m, typ.Method(m).Index)
		fmt.Println("m", m, typ.Method(m).PkgPath)
	}
}
