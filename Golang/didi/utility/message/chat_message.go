package message

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

const (
	CHATMSG = 2
)

type ChatMsg struct {
	Addr    string
	Content string
}

func (msg ChatMsg) String() string {
	return msg.Addr + ": " + msg.Content
}

func (msg ChatMsg) Serialize() ([]byte, error) {
	msgBuf, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, msg.MessageNumber())
	binary.Write(buf, binary.BigEndian, int32(len(msgBuf)))
	buf.Write(msgBuf)

	return buf.Bytes(), err
}

func (msg ChatMsg) MessageNumber() int32 {
	return CHATMSG
}

func DeserializeChatMsg(buf []byte) (Message, error) {
	msg := new(ChatMsg)
	err := json.Unmarshal(buf, msg)
	if err != nil {
		return nil, err
	} else {
		return *msg, err
	}
}
