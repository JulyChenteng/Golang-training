package main

import (
	"fmt"
	"net"
	"syscall"

	"./msghandler"
	"./service"
	"os"
	"os/signal"

	"../utility/users"
)

const (
	SERVERADDR = "127.0.0.1:5000"

	MAXCONNS = 1024
)

func main() {
	listener := service.InitServer(SERVERADDR)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error : ", err)
			continue
		}

		if service.ConnNum() < MAXCONNS {
			switch users.CheckClientType(conn.RemoteAddr().String()) {
			case users.PASSENGER:
				service.AddConn(conn, users.PASSENGER)
				go passengerConnHandler(conn)
			case users.OWNER:
				service.AddConn(conn, users.OWNER)
				go ownerConnHandler(conn)
			default:
				conn.Close()
				fmt.Println("Illegal user!")
			}
		} else {
			conn.Close()
		}
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c

		service.Pool.Shutdown()
		listener.Close()
		service.Stop()
	}()
}

func passengerConnHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr()
	fmt.Println("Connect from: [Passenger] ", connFrom)

	for {
		var buf = make([]byte, msghandler.MAXREAD)
		len, err := conn.Read(buf[0:])

		switch err {
		case nil:
			msghandler.HandleMsg(conn, buf, len)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	service.RemoveConn(conn)

	err := conn.Close()
	fmt.Println("Closed connection: ", connFrom)
	service.CheckError(err, "Close")
}

func ownerConnHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr()
	fmt.Println("Connect from: [Owner] ", connFrom)

	//将所有未被处理的订单信息发送车主
	service.SendUnProMsgToOwner(conn)
	for {
		var buf = make([]byte, msghandler.MAXREAD)
		len, err := conn.Read(buf[0:])

		switch err {
		case nil:
			msghandler.HandleMsg(conn, buf, len)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	service.RemoveConn(conn)
	err := conn.Close()
	fmt.Println("Closed connection: ", connFrom)

	service.CheckError(err, "Close")
}
