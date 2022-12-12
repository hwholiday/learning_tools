package main

import "fmt"

func main() {
	user, err := NewUserBuilder().WithName("").Builder()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user)
}
