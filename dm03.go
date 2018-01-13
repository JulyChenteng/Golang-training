package main

import (
	"fmt"
)

/*
const a int = 1
const b = 'A'

const (
	c, d, e = 2, 3, 4
)

const (
	f = 5
	g
	h
)

const (
	j = "aaa"
	k = len(j)
	l
)

const (
	m, n = 66, 77
	o, p
)

func main() {
	fmt.Println(a) //1
	fmt.Println(b) //65
	fmt.Println(c) //2
	fmt.Println(d) //3
	fmt.Println(e) //4
	fmt.Println(f) //5
	fmt.Println(g) //5
	fmt.Println(h) //5
	fmt.Println(j) //aaa
	fmt.Println(k) //3
	fmt.Println(l) //3
	fmt.Println(m) //66
	fmt.Println(n) //77
	fmt.Println(o) //66
	fmt.Println(p) //77
}
*/

const (
	a = 'A'
	b
	c = iota
	d
)

const (
	e = iota
	f
	g
)

func main() {
	fmt.Println(a) //65
	fmt.Println(b) //65
	fmt.Println(c) //2
	fmt.Println(d) //3

	fmt.Println(e) //0
	fmt.Println(f) //1
	fmt.Println(g) //2

	/////////////////////运算符练习////////////////////////
	/*
		6: 0110
		11: 1011
		& 0010 = 2
		| 1111 = 15
		^ 1101 = 13
		&^ 0100 = 4 第二个数值为1时置零
	*/
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

	x := 0
	if x > 0 && (10/x) > 1 {
		fmt.Println("OK")
	}
}
