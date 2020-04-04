package msghandler

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"../../../utility/message"
	"../../service"
	"time"
)

const (
	MAXREAD = message.MAXPKGSIZE
)

var (
	handlerMaps = make(service.HandlerMaps)
	OrderList   = make([]message.OrderMsg, 0, 20)
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
			PrintLocalOrders()

			if len(OrderList) != 0 {
				SendOrderMsgByMsgNo(conn)
			}
		}*/
	}
}

func ProcessOrderMsg(conn net.Conn, info message.Message) {
	fmt.Println("Process OrderMsg from Server......")

	var mutex sync.Mutex
	reader := bufio.NewReader(os.Stdin)

	if msg, ok := info.(message.OrderMsg); ok {
		if strings.Compare(msg.Owner, conn.LocalAddr().String()) == 0 && msg.State == message.ORDER_PROCSSED {
			fmt.Println(msg.Owner, "抢单成功，是否和乘客聊天?(Y/N):")

			for index, value := range OrderList {
				if value.Id == msg.Id {
					OrderList = append(OrderList[0:index], OrderList[index+1:]...)
				}
			}

			input, _ := reader.ReadString('\n')
			str := strings.Trim(input, "\r\n")

			if str == "Y" {
				service.ClearTerminal()
				SendChatMsg(conn, msg.Passenger) //发送聊天信息
			} else {
				service.ClearTerminal()
				//打印本地所有订单信息
				PrintLocalOrders()

				if len(OrderList) != 0 {
					SendOrderMsgByMsgNo(conn)
				}
			}

			return
		}

		if msg.Passenger == "" {
			fmt.Println("该单已被抢...")
			//从本地列表中删除该订单
			for index, value := range OrderList {
				if value.Id == msg.Id {
					OrderList = append(OrderList[0:index], OrderList[index+1:]...)
				}
			}

			PrintLocalOrders()
			if len(OrderList) != 0 {
				SendOrderMsgByMsgNo(conn)
			}
		} else {
			mutex.Lock()
			OrderList = append(OrderList, msg)
			mutex.Unlock()

			PrintLocalOrders()
			SendOrderMsgByMsgNo(conn)
		}
	}
}

func SendOrderMsgByMsgNo(conn net.Conn) {
	for {
		var msgId = -1
		fmt.Println("抢单请输入单号,刷新按r,取消q: ")
		var str string
		fmt.Scanf("%s", &str)
		if str == "r" {
			service.ClearTerminal()
			PrintLocalOrders()
			continue
		} else if str == "q" {
			return
		} else {
			var err error
			msgId, err = strconv.Atoi(str)
			if err != nil {
				msgId = -1
			}
		}

		var msg message.OrderMsg
		var isExist = false //用于判断单号是否合法
		for _, value := range OrderList {
			if value.Id == msgId {
				msg = value
				isExist = true
			}
		}

		if isExist {
			msg.Owner = conn.LocalAddr().String()
			msg.State = message.ORDER_PROCSSED

			buf, _ := msg.Serialize()
			conn.Write(buf)

			return
		} else {
			fmt.Println("Order Id error!")
		}
	}
}

func HandleMsg(conn net.Conn, buf []byte, len int) error {
	var msgType, length int32

	//获取消息类型
	typeBuf := bytes.NewBuffer(buf[0:message.MSGTYPE])
	err := binary.Read(typeBuf, binary.BigEndian, &msgType)
	if err != nil {
		return err
	}

	//获取消息长度
	lenBuf := bytes.NewBuffer(buf[message.MSGTYPE : message.MSGTYPE+message.MSGLENGTH])
	err = binary.Read(lenBuf, binary.BigEndian, &length)
	if err != nil {
		return nil
	}

	if length > message.MSGMAXLEN {
		fmt.Println("Message's length overflow!")
		return err
	}

	//获取消息内容
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
	PrintLocalOrders()

	if len(OrderList) != 0 {
		SendOrderMsgByMsgNo(conn)
	}
}

func PrintLocalOrders() {
	//打印本地所有订单信息
	if len(OrderList) != 0 {
		fmt.Println("单号\t乘客\t\t起点\t目的地")
		for _, v := range OrderList {
			fmt.Printf("%d\t%s\t   %s\t%s\n", v.Id, v.Passenger, v.Origin, v.Destination)
		}
	} else {
		fmt.Println("暂无订单信息，等待中... ")
	}
}

func SendHBMsg(conn net.Conn) {
	for {
		msg := message.HBMsg{}

		buf, _ := msg.Serialize()
		conn.Write(buf)

		time.Sleep(10 * time.Second)
	}
}
