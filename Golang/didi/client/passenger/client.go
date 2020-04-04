package main

import (
	"../service"
	"./msghandler"
	"flag"
	"fmt"
	"syscall"
)

const (
	SERVERADDR = "127.0.0.1:5000"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage : host port")
	}

	hostAndPort := flag.Arg(0) + ":" + flag.Arg(1)

	dialer := service.InitClient(hostAndPort)
	conn, err := dialer.Dial("tcp", SERVERADDR)
	service.CheckError(err, "Dial : ")
	defer conn.Close()

	fmt.Println("On connect")
	go msghandler.SendHBMsg(conn)
	msghandler.SendOrderMsg(conn)
	for {
		buf := make([]byte, msghandler.MAXREAD)
		len, err := conn.Read(buf)

		switch err {
		case nil:
			go msghandler.HandleMsg(conn, buf, len)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err = conn.Close()
	fmt.Println("Server Closed connection: ", conn.RemoteAddr())
	service.CheckError(err, "Close: ")
}
