package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	Name string
	time.Time
}

func (m MyError) Error() string {
	return fmt.Sprintf("%v %v", m.Name, m.Time)
}

func TestErr() error {
	return MyError{
		Name: "AAAA",
		Time: time.Now(),
	}
}

func main() {

	err := errors.New("test")
	fmt.Println(err)
	err = fmt.Errorf("%s %w",err.Error(), TestErr())
	fmt.Println(err)
	err = fmt.Errorf("2 %w", err)
	fmt.Println(err.Error())
	var myErr MyError
	if errors.As(err, &myErr) {
		fmt.Println(myErr)
	}
	fmt.Println(errors.Unwrap(err))
	fmt.Println(errors.Is(err, errors.Unwrap(err)))
}
