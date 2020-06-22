package msg

import (
	"learning_tools/websocket/pb"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	NewMsgProtocol(true)
	os.Exit(m.Run())
}

func TestNewMsgProtocol(t *testing.T) {
	p := GetMsgProtocol()
	err := p.Register(&pb.Ping{}, 1)
	if err != nil {
		t.Error(err)
	}
	data, err := p.Marshal(&pb.Ping{Times: 1})
	if err != nil {
		t.Error(err)
	}
	info, err := p.Unmarshal(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(info.(*pb.Ping))
}
