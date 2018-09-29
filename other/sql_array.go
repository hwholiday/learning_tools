package utils

import (
	"strings"
	"fmt"
)

func Int32ToStringArray(in []int32) (out []string) {
	for _, v := range in {
		out = append(out, fmt.Sprint(v))
	}
	return
}

func AddString(s string, v []string) (string) {
	if len(s) <= 0 { //第一次添加
		s = strings.Join(v, ",")
	} else {
		s += strings.Join(v, ",")
	}
	return s
}

func DelString(s, v string) (string) {
	var index int
	var ok bool
	data := removeRep(strings.Split(s, ","))
	for k, val := range data {
		if val == v {
			index = k
			ok = true
			break
		}
	}
	if ok {
		data = append(data[:index], data[index+1:]...)
	}
	return strings.Join(data, ",")
}

//去除数字里面重复的
func removeRep(slc []string) []string {
	var result []string
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}
