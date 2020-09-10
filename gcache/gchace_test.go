package gcache

import "testing"

func TestNewGroups(t *testing.T) {
	g := NewGroups("user", 0, GetterFunc(
		func(key string) ([]byte, error) {
			t.Log("GetterFunc key : ", key)
			if key == "1" {
				return []byte("111111111111111"), nil
			}
			return nil, nil
		}))
	val, err := g.Get("1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(val))
	val, err = g.Get("1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(val))
}
