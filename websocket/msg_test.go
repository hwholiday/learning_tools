package gateway

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	NewMsgProtocol(true)
	os.Exit(m.Run())
}

func TestNewMsgProtocol(t *testing.T) {
	p := GetMsgProtocol()
	err := p.Register(&Ping{}, 1)
	if err != nil {
		t.Error(err)
	}
	data, err := p.Marshal(&Ping{Times: 1})
	if err != nil {
		t.Error(err)
	}
	info, err := p.Unmarshal(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(info.(*Ping))
}
