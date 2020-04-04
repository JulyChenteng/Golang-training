package defs

import (
	"net"

	"../../utility/users"
)

type Conn struct {
	Conn    net.Conn
	CliType int32
}

func (conn Conn) String() string {
	str := conn.RemoteAddr()
	if conn.CliType == users.PASSENGER {
		return "[PASSENGER:" + str + "]"
	} else if conn.CliType == users.OWNER {
		return "[OWNER:" + str + "]"
	}

	return "error Addr : " + str
}

func (conn *Conn) Write(buf []byte) (int, error) {
	return conn.Conn.Write(buf)
}

func (conn *Conn) Read(buf []byte) (int, error) {
	return conn.Conn.Read(buf)
}

func (conn *Conn) LocalAddr() string {
	return conn.Conn.LocalAddr().String()
}

func (conn *Conn) RemoteAddr() string {
	return conn.Conn.RemoteAddr().String()
}

func (conn *Conn) Close() error {
	return conn.Conn.Close()
}
