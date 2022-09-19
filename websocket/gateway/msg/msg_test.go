package msg

import (
	"os"
	"testing"

	"github.com/hwholiday/learning_tools/websocket/pb"
)

func TestMain(m *testing.M) {
	NewMsgProtocol(true)
	os.Exit(m.Run())
}

func TestNewMsgProtocol(t *testing.T) {
	p := GetMsgProtocol()
	p.Register(&pb.Ping{}, 1)
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
