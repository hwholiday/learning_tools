package main

import (
	"testing"
)

func TestBinaryFind(t *testing.T) {
	var sortedArray []int = []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	t.Log(BinaryFind(sortedArray, 6))
}

func BinaryFind(array []int, find int) (val int) {
	var (
		leaf  = 0
		right = len(array)
	)
	for leaf <= right {
		mod := (leaf + right) / 2
		modVal := array[mod]
		if modVal > find {
			right = mod - 1
		} else if modVal < find {
			leaf = mod + 1
		} else {
			return mod
		}
	}
	return -1
}
