package msghandler

import (
	"fmt"

	"../../utility/message"
	"../../utility/users"
	"../defs"
	"../service"
	"bytes"
	"encoding/binary"
	"net"
	"sync"
	"time"
)

const (
	MAXREAD = message.MAXPKGSIZE

	TIMEOUT = 10
)

func init() {
	service.Register(message.CHATMSG, message.DeserializeChatMsg, ProcessChatMsg)
	service.Register(message.ORDERMSG, message.DeserializeOrderMsg, ProcessOrderMsg)
}

func ProcessChatMsg(conn defs.Conn, info message.Message) {
	fmt.Println("Process ", conn, "'s ChatMsg")
	fmt.Println(info)

	if msg, ok := info.(message.ChatMsg); ok {
		addr := msg.Addr
		msg.Addr = conn.RemoteAddr()
		fmt.Println(msg)
		//如果是乘客发送的信息，检查车主是否online
		if conn.CliType == users.PASSENGER {
			if service.IsOwnerOnline(addr) {
				c := service.GetConn(addr)
				buf, _ := msg.Serialize()
				c.Write(buf)
			}
		} else if conn.CliType == users.OWNER {
			if service.IsPassengerOnline(addr) {
				c := service.GetConn(addr)
				buf, _ := msg.Serialize()
				c.Write(buf)
			}
		}
	}
}

func ProcessOrderMsg(conn defs.Conn, info message.Message) {
	fmt.Println("Process ", conn, "'s OrderMsg")
	var mutex sync.Mutex

	if msg, ok := info.(message.OrderMsg); ok {
		//处理乘客下单信息
		if conn.CliType == users.PASSENGER {
			//将订单保存至订单列表中
			mutex.Lock()
			msg.Id = len(service.OrderList)
			msg.Passenger = conn.RemoteAddr()
			service.OrderList = append(service.OrderList, msg)
			mutex.Unlock()

			//将新的订单信息发送给所有车主
			buf, _ := msg.Serialize()

			for _, c := range service.Conns {
				if c.CliType == users.OWNER {
					c.Write(buf)
				}
			}
		} else if conn.CliType == users.OWNER { //处理车主接单后发送的信息
			var isWrited = false //用于判断服务器接单信息是否写入成功

			//更改服务器中保存的订单信息
			for index, value := range service.OrderList {
				if value.Id == msg.Id {
					mutex.Lock()
					if value.State == message.ORDER_NOPROCESS {
						service.OrderList[index] = msg
						isWrited = true
					}
					mutex.Unlock()
				}
			}
			fmt.Println(msg.Id)
			fmt.Println(service.OrderList)
			if isWrited {
				buf, _ := msg.Serialize()
				//通知乘客
				addr := msg.Passenger //获取乘客信息
				if service.IsPassengerOnline(addr) {
					c := service.GetConn(addr)
					c.Write(buf)
				}
				conn.Write(buf) //通知车主成功接单
			} else {
				//通知车主 抢单失败
				msg = message.OrderMsg{Id: msg.Id}
				buf, _ := msg.Serialize()
				conn.Write(buf)
			}
		}
	}
}

func HandleMsg(conn net.Conn, buf []byte, len int) error {
	var msgType, length int32

	//	isHBMsg := make(chan bool)
	//	go ProcessHBMsg(conn, isHBMsg, TIMEOUT)

	typeBuf := bytes.NewBuffer(buf[0:message.MSGTYPE])
	err := binary.Read(typeBuf, binary.BigEndian, &msgType)
	if err != nil {
		return err
	}

	//处理心跳消息
	if msgType == 0 {
		//isHBMsg <- true
		fmt.Println("Process HBMsg from client: ", conn.RemoteAddr().String())
		conn.SetDeadline(time.Now().Add(time.Duration(TIMEOUT) * time.Second))
		return nil
	}

	lenBuf := bytes.NewBuffer(buf[message.MSGTYPE : message.MSGTYPE+message.MSGLENGTH])
	err = binary.Read(lenBuf, binary.BigEndian, &length)
	if err != nil {
		return nil
	}

	if length > message.MSGMAXLEN {
		fmt.Println("Message's length overflow!")
		return err
	}

	msgBuf := buf[message.MSGTYPE+message.MSGLENGTH : len]
	deserialize := service.GetMsgDeserializeFunc(msgType)
	msg, err := deserialize(msgBuf)
	if err != nil {
		fmt.Println("Message Deserialize failed!")
		return err
	}

	c := service.GetConn(conn.RemoteAddr().String())
	handler := service.GetMsgHandler(msgType)

	service.Pool.Run(service.Job{
		Data: msg,
		Conn: c,
		Proc: handler,
	})

	return nil
}

//func ProcessHBMsg(conn net.Conn, ch chan bool, timeout time.Duration) {
//	fmt.Println("Process HBMsgFrom ", conn)
//	select {
//	case _ = <-ch:
//		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
//		break
//	case <-time.After(time.Second * timeout):
//		fmt.Println("timeout")
//		service.RemoveConn(conn)
//		conn.Close()
//	}
//}
