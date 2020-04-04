package message

import (
	"encoding/json"

	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	ORDER_PROCSSED  = true
	ORDER_NOPROCESS = false

	ORDERMSG = 1
)

type OrderMsg struct {
	Id          int
	Passenger   string
	Owner       string
	Origin      string //起点
	Destination string //目的地
	State       bool   //状态
	Completed   bool
}

func (order OrderMsg) String() string {
	return fmt.Sprintf("%d\t%s\t%s\t%s\t%s", order.Id, order.Passenger, order.Owner, order.Origin, order.Destination)
}

func (order OrderMsg) Serialize() ([]byte, error) {
	msgBuf, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, order.MessageNumber())
	binary.Write(buf, binary.BigEndian, int32(len(msgBuf)))
	buf.Write(msgBuf)

	return buf.Bytes(), err
}

func (order OrderMsg) MessageNumber() int32 {
	return ORDERMSG
}

func DeserializeOrderMsg(buf []byte) (Message, error) {
	order := new(OrderMsg)
	err := json.Unmarshal(buf, &order)
	if err != nil {
		return nil, err
	} else {
		return *order, nil
	}
}
