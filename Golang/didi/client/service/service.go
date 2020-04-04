package service

import (
	"net"

	"../../utility/message"
	"os"
	"os/exec"
)

type Handler func(net.Conn, message.Message)
type DeserializeFunc func([]byte) (message.Message, error)

type HandlerAndFunc struct {
	Handle      Handler
	Deserialize DeserializeFunc
}

type HandlerMaps map[int32]HandlerAndFunc

func InitClient(hostAndPort string) net.Dialer {
	clientAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	CheckError(err, "Resolve TCPAddr: ")

	return net.Dialer{LocalAddr: clientAddr}
}

func CheckError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}

func Register(handlerMaps HandlerMaps, msgType int32, f DeserializeFunc, handler Handler) {
	handlerMaps[msgType] = HandlerAndFunc{Deserialize: f, Handle: handler}
}

func GetMsgHandler(handlerMaps HandlerMaps, msgType int32) Handler {
	hf, ok := handlerMaps[msgType]
	if !ok {
		return nil
	}

	return hf.Handle
}

func GetMsgDeserializeFunc(handlerMaps HandlerMaps, msgType int32) DeserializeFunc {
	hf, ok := handlerMaps[msgType]
	if !ok {
		return nil
	}

	return hf.Deserialize
}

//清空终端
func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
