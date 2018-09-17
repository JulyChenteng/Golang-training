package main

import (
	"fmt"
	"strconv"
)

var i int = 10

func main() {
	str := "112"
	i := 56
	var j int64
	j = 667

	//string到int
	num, _ := strconv.Atoi(str)
	//string到int64
	num64, _ := strconv.ParseInt(str, 10, 64)
	//int到string
	str1 := strconv.Itoa(i)
	str2 := fmt.Sprintf("%d", i) // 通过Sprintf方法转换
	//int64到string
	str3 := strconv.FormatInt(j, 10)

	fmt.Println(num)
	fmt.Println(num64)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
