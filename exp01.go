package main

import (
	"fmt"
)

//如果要传任意类型，需要指定类型为interface{}
func MyPrintf1(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 type.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

//...type本质上是一个数组切片，也就是[]type
func MyPrintf2(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello world!"
	var v4 float32 = 1.234

	MyPrintf1(v1, v2, v3, v4)
	MyPrintf2(1, 2, 3, 4, 5)
}
