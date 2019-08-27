package LRU

import "testing"

func TestNewList(t *testing.T) {
	l := NewLru(2)
	t.Log(l.Add("1", "1"))
	t.Log(l.Add("2", "2"))
	for _, v := range l.GetAll() {
		t.Log("Key", v.Key, "Val", v.Val)
	}
	t.Log(l.Add("3", "3"))
	t.Log(l.Add("4", "4"))
	for _, v := range l.GetAll() {
		t.Log("Key", v.Key, "Val", v.Val)
	}
}
