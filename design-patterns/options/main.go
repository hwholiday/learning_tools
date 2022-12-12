package main

import "fmt"

func main() {
	opt := NewUserOptions(WithName("aaa"), WithAge(2))
	fmt.Printf("%+v", opt)
}
