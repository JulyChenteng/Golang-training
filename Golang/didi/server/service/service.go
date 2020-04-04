package service

import (
	"fmt"
	"net"
	"strings"

	"../../utility/message"
	"../../utility/users"
	"../defs"
	"sync"
)

type Handler func(defs.Conn, message.Message)
type DeserializeFunc func([]byte) (message.Message, error)

type HandlerAndFunc struct {
	Handle      Handler
	Deserialize DeserializeFunc
}

var (
	Conns       = make([]defs.Conn, 0, 20)
	HandlerMaps = make(map[int32]HandlerAndFunc)

	OrderList = make([]message.OrderMsg, 0, 20)

	mutex sync.Mutex
)

func InitServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	CheckError(err, "Resolving address:port failed: '"+hostAndPort+"'")

	listener, err := net.ListenTCP("tcp", serverAddr)
	CheckError(err, "ListenTCP")

	fmt.Println("listening to : ", listener.Addr())

	return listener
}

func CheckError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}

func IsOwnerOnline(addr string) bool {
	for _, v := range Conns {
		if strings.Compare(addr, v.RemoteAddr()) == 0 && users.OWNER == v.CliType {
			return true
		}
	}

	return false
}

func IsPassengerOnline(addr string) bool {
	for _, v := range Conns {
		if strings.Compare(addr, v.RemoteAddr()) == 0 && users.PASSENGER == v.CliType {
			return true
		}
	}

	return false
}

func GetConn(addr string) defs.Conn {
	mutex.Lock()
	for _, v := range Conns {
		if strings.Compare(addr, v.RemoteAddr()) == 0 {
			return v
		}
	}
	mutex.Unlock()

	return defs.Conn{Conn: nil, CliType: users.UNKNOW}
}

func RemoveConn(conn net.Conn) {
	for index, val := range Conns {
		if strings.Compare(conn.RemoteAddr().String(), val.RemoteAddr()) == 0 {
			Conns = append(Conns[0:index], Conns[index+1:]...)
			return
		}
	}
}

func Register(msgType int32, f DeserializeFunc, handler func(conn defs.Conn, message message.Message)) {
	HandlerMaps[msgType] = HandlerAndFunc{Deserialize: f, Handle: handler}
}

func GetMsgHandler(msgType int32) Handler {
	hf, ok := HandlerMaps[msgType]
	if !ok {
		return nil
	}

	return hf.Handle
}

func GetMsgDeserializeFunc(msgType int32) DeserializeFunc {
	hf, ok := HandlerMaps[msgType]
	if !ok {
		return nil
	}

	return hf.Deserialize
}

func SendUnProMsgToOwner(conn net.Conn) {
	for _, value := range OrderList {
		if value.State == message.ORDER_NOPROCESS {
			buf, _ := value.Serialize()
			conn.Write(buf)
		}
	}
}

func AddConn(conn net.Conn, cliType int32) {
	Conns = append(Conns, defs.Conn{Conn: conn, CliType: cliType})
}

func Stop() {
	for _, conn := range Conns {
		conn.Conn.Close()
		fmt.Println("close client : ", conn)
	}
}

func ConnNum() int {
	return len(Conns)
}
