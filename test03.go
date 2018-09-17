package main

import (
	"fmt"
)

/*
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
*/

func main() {

	a := make([]int, 0, 6)        //
	b := append(a, 1, 2, 3, 4, 5) //前五个元素，未超过a的大小，与a共用一块内存
	c := append(a[1:3], 8, 9, 10) //后五个元素

	fmt.Println(a) //[]  ？？？？？？
	fmt.Println(b) //[1 2 3 8 9]
	fmt.Println(c) //[2 3 8 9 10]

	/*
		a := []int{1, 2, 3, 4, 5}
		b := append(a[1:3], 8, 9, 10)  //超过a的容量，重新分配内存

		fmt.Println(a) //[1 2 3 4 5]
		fmt.Println(b) //[2 3 8 9 10]
	*/

	/*
		a := []int{1, 2, 3, 4, 5}
		b := append(a[1:3], 8, 9)

		fmt.Println(a) //[1 2 3 8 9]
		fmt.Println(b) //[2 3 8 9]
	*/
}
