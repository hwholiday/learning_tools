package cache

import (
	"errors"
	"regexp"
	"strconv"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var ErrParameter = errors.New("parameter error")
var ErrUnitUndefined = errors.New("unit undefined")
var ErrNotEnoughSpace = errors.New("not enough space")

var SizeMap = map[string]int{
	"KB": KB,
	"MB": MB,
	"GB": GB,
}

type Object struct {
	Data       interface{}
	Inquire    int
	Size       int
	Expiration int64
}

type node struct {
	key  string
	size int
	num  int
}

func GetSize(size string) (num int, err error) {
	compile, err := regexp.Compile("^[1-9]\\d*|[A-Z]+")
	if err != nil {
		return 0, err
	}
	result := compile.FindAllString(size, -1)
	if len(result) != 2 {
		return 0, ErrParameter
	}
	number, err := strconv.Atoi(result[0])
	if err != nil {
		return 0, err
	}
	unit, ok := SizeMap[result[1]]
	if !ok {
		return 0, ErrUnitUndefined
	}
	return number * unit, nil
}
