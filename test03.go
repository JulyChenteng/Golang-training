package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	GB
	TB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(GB)
	fmt.Println(TB)
}
