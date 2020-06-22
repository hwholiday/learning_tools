package msg

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"math"
	"reflect"
)

var mgPrt *MsgProtocol

//协议
// id + proto.Message
type MsgProtocol struct {
	msgID        map[reflect.Type]uint16
	msgInfo      map[uint16]reflect.Type
	useBigEndian bool
}

func NewMsgProtocol(useBigEndian bool) {
	mgPrt = &MsgProtocol{
		msgID:        make(map[reflect.Type]uint16),
		msgInfo:      make(map[uint16]reflect.Type),
		useBigEndian: useBigEndian,
	}
}

func GetMsgProtocol() *MsgProtocol {
	if mgPrt == nil {
		panic("msg prt nil")
	}
	return mgPrt
}

func (m *MsgProtocol) Register(msg proto.Message, eventType uint16) error {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return ErrMsgNotProto
	}
	if len(m.msgInfo) >= math.MaxUint16 {
		return ErrProtocol
	}
	m.msgInfo[eventType] = msgType
	m.msgID[msgType] = eventType
	return nil
}

func (m *MsgProtocol) Marshal(msg interface{}) ([]byte, error) {
	msgType := reflect.TypeOf(msg)
	event, ok := m.msgID[msgType]
	if !ok {
		return nil, ErrNotRegister
	}
	data, err := proto.Marshal(msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	var (
		id      = make([]byte, 2)
		ptrData = make([]byte, 2+len(data))
	)
	if m.useBigEndian {
		binary.BigEndian.PutUint16(id, event)
	} else {
		binary.LittleEndian.PutUint16(id, event)
	}
	copy(ptrData[:2], id)
	copy(ptrData[2:], data)
	return ptrData, nil
}

func (m *MsgProtocol) Unmarshal(msg []byte) (interface{}, error) {
	if len(msg) < 2 {
		return nil, ErrMsgShort
	}
	var id uint16
	if m.useBigEndian {
		id = binary.BigEndian.Uint16(msg[:2])
	} else {
		id = binary.LittleEndian.Uint16(msg[:2])
	}
	msgType, ok := m.msgInfo[id]
	if !ok {
		return nil, ErrNotRegister
	}
	var data = reflect.New(msgType.Elem()).Interface()
	err := proto.Unmarshal(msg[2:], data.(proto.Message))
	if err != nil {
		return nil, err
	}
	return data, nil
}
