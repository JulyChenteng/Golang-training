package message

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
)

const (
	HBMSG = 0
)

type HBMsg struct{}

func (msg HBMsg) MessageNumber() int32 {
	return HBMSG
}

func (msg HBMsg) Serialize() ([]byte, error) {
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
