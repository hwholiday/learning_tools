package tool

import (
	"github.com/google/uuid"
	"net/url"
	"strings"
	"testing"
)

func TestDncrypt(t *testing.T) {
	//"1837032330871439"
	str, err := AesECBEncrypt("1837032330", []byte("123123123"))
	if err != nil {
		panic(err)
	}
	t.Log(str)
	res, err := AesECBDecrypt("1837032330", str)
	if err != nil {
		panic(err)
	}
	t.Log(string(res))
}

func TestUUID(t *testing.T) {
	t.Log(strings.ReplaceAll(uuid.New().String(), "-", ""))
}

func TestUrl(t *testing.T) {
	URL := "http://127.0.0.1:8888/a/b"
	u, err := url.Parse(URL)
	if err != nil {
		t.Log(err)
	}
	t.Log(u.Host)
}
