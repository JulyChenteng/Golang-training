package msghandler

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"

	"../../../utility/message"
	"../../service"
	"time"
)

const (
	MAXREAD = message.MAXPKGSIZE
)

var (
	handlerMaps = make(service.HandlerMaps)
)

func init() {
	service.Register(handlerMaps, message.ORDERMSG, message.DeserializeOrderMsg, ProcessOrderMsg)
	service.Register(handlerMaps, message.CHATMSG, message.DeserializeChatMsg, ProcessChatMsg)
}

func ProcessChatMsg(conn net.Conn, info message.Message) {
	fmt.Println()
	fmt.Println("Process ChatMsg from Server......")

	if msg, ok := info.(message.ChatMsg); ok {
		fmt.Println(msg.Addr, ": ", msg.Content) //打印接收的聊天信息

		/*if msg.Content == "bye" {
			time.Sleep(2 * time.Second)
			service.ClearTerminal()
			SendOrderMsg(conn)
		}*/
	}
}

//处理订单信息，乘客端收到的订单信息是已被接单的订单信息
func ProcessOrderMsg(conn net.Conn, info message.Message) {
	fmt.Println("Process OrderMsg from Server......")
	reader := bufio.NewReader(os.Stdin)

	if msg, ok := info.(message.OrderMsg); ok {
		fmt.Println(msg.Owner, "已接单，是否和车主聊天?(Y/N):")

		input, _ := reader.ReadString('\n')
		str := strings.Trim(input, "\r\n")

		if str == "Y" {
			service.ClearTerminal()
			SendChatMsg(conn, msg.Owner)
		} else {
			service.ClearTerminal()
			SendOrderMsg(conn)
		}
	}
}

func HandleMsg(conn net.Conn, buf []byte, len int) error {
	var msgType, length int32

	typeBuf := bytes.NewBuffer(buf[0:message.MSGTYPE])
	err := binary.Read(typeBuf, binary.BigEndian, &msgType)
	if err != nil {
		return err
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
	deserialize := service.GetMsgDeserializeFunc(handlerMaps, msgType)
	msg, err := deserialize(msgBuf)
	if err != nil {
		fmt.Println("Message Deserialize failed!")
		return err
	}

	handler := service.GetMsgHandler(handlerMaps, msgType)
	handler(conn, msg)

	return nil
}

//发送下单信息
func SendOrderMsg(conn net.Conn) {
	fmt.Println("请下单，请输入origin and destination:")

	var origin, destination string
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		str := strings.Trim(input, "\r\n")
		fmt.Sscanf(str, "%s %s", &origin, &destination)

		if origin == "" || destination == "" {
			fmt.Println("Origin or destination Error!")
			fmt.Println("Please input again: ")
		} else {
			break
		}
	}

	msg := message.OrderMsg{
		Origin:      origin,
		Destination: destination,
	}
	msg.Passenger = conn.LocalAddr().String()

	buff, _ := msg.Serialize()
	conn.Write(buff)
}

//发送聊天信息
func SendChatMsg(conn net.Conn, addr string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(conn.LocalAddr(), ": ")
		input, _ := reader.ReadString('\n')
		str := strings.Trim(input, "\r\n")

		//打包聊天信息
		newMsg := message.ChatMsg{
			Addr:    addr,
			Content: str,
		}

		buf, _ := newMsg.Serialize()
		conn.Write(buf)

		if str == "bye" {
			break
		}
	}

	service.ClearTerminal()
	SendOrderMsg(conn)
}

func SendHBMsg(conn net.Conn) {
	for {
		msg := message.HBMsg{}

		buf, _ := msg.Serialize()
		conn.Write(buf)

		time.Sleep(10 * time.Second)
	}
}
