package main

import (
	"fmt"
)

func main() {
	xx := [2]int{1} //第一个元素为1，其他默认为0
	fmt.Println(xx)
	a := [20]int{19: 1} //a[19]为1，其余默认为0
	fmt.Println(a)
	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	c := [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	fmt.Println(c)
	d := [...]int{29: 1}
	fmt.Println(d)
	var p *[30]int = &d //指向包含30个整型的数组（数组指针）
	fmt.Println(p)
	x, y := 1, 2
	q := [...]*int{&x, &y} //指针数组，数组中的元素均为指针类型
	fmt.Println(q)
	// 指向数组的指针如何写?
	// r := [...]*[...]int{&c, &d} //error use &c (type *[3]int) as type *[...]int in array or slice literal
	r := [...]*[5]int{&b, &c} // b和c数组都是相同类型的数组(长度一样)
	fmt.Println(r)

	aa := new([10]int) // 指向数组的指针 new返回一个指针
	aa[1] = 2
	fmt.Println(aa)
	bb := [10]int{}
	bb[1] = 2
	fmt.Println(bb)

	aaa := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}} // 注意,最后一个}不能单独占一行
	fmt.Println(aaa)

	bbb := [2][3]int{
		{1: 1},
		{2: 2}} // 注意,最后一个}不能单独占一行
	fmt.Println(bbb)

	ccc := [...][3]int{ // NOTE: 第二维不能用...省略
		{1: 1},
		{2: 2}} // 注意,最后一个}不能单独占一行
	fmt.Println(ccc)

	ppp := [...]int{5, 2, 6, 3, 9}
	fmt.Println(ppp)
	cnt := len(ppp)
	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			if ppp[i] < ppp[j] {
				temp := ppp[i]
				ppp[i] = ppp[j]
				ppp[j] = temp
			}
		}
	}
	fmt.Println(ppp)
}
