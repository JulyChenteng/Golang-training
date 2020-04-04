package users

import (
	"strconv"
)

const (
	PASSENGER = 1
	OWNER     = 2
	UNKNOW    = -1
)

var (
	Passengers = make([]string, 0, 10)
	Owners     = make([]string, 0, 10)
)

func init() {
	initOwners()
	initPassengers()
}

func initPassengers() {
	for i := 6000; i < 6005; i++ {
		addr := "127.0.0.1:" + strconv.Itoa(i)
		Passengers = append(Passengers, addr)
	}

	for i := 6000; i < 6005; i++ {
		addr := "169.254.46.227" + strconv.Itoa(i)
		Passengers = append(Passengers, addr)
	}
}

func initOwners() {
	for i := 7000; i < 7005; i++ {
		addr := "127.0.0.1:" + strconv.Itoa(i)
		Owners = append(Owners, addr)
	}
}

//检测用户的合法性，返回1表示为乘客，2表示车主，-1表示用户不合法
func CheckClientType(addr string) int {
	for _, v := range Passengers {
		if addr == v {
			return PASSENGER
		}
	}

	for _, v := range Owners {
		if addr == v {
			return OWNER
		}
	}

	return UNKNOW
}
