package main

import (
	"fmt"
	"testing"
)

type Num interface {
	int64 | float64
}

func TestT(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	fmt.Println(SumIntsOrFloats[string, int64](ints))
	fmt.Println(SumNum(ints))
}
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
func SumNum[K comparable, V Num](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
