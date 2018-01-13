package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		var m1 map[int]string = make(map[int]string)
		m1[1] = "OK"
		fmt.Println(m1[1])

		delete(m1, 1)
		fmt.Println(m1[1])

		m2 := make(map[int]map[int]string)

		a, ok := m2[1][1] //多返回值，如果m2[1][1]存在，则返回其值和bool类型，存在为true，否则为false
		if !ok {
			m2[1] = make(map[int]string)
		}

		m2[1][1] = "GOOD"
		a = m2[1][1]

		fmt.Println(a, ok)
	*/

	/*
		//迭代, i为slice的索引，v为slice中值的拷贝
		for i, v := range slice名 {

		}
	*/

	/*
		sm := make([]map[int]string, 5)
		//同样v为一份拷贝，对v的操作对原map不会产生影响
		for _, v := range sm {
			v = make(map[int]string, 1)
			v[1] = "OK"
			fmt.Println(v)
		}
		fmt.Println(sm)

		//要想修改map的值，则应该通过索引操作
		for k := range sm {
			sm[k] = make(map[int]string, 1)
			sm[k][1] = "OK"
			fmt.Println(sm[k])
		}
		fmt.Println(sm)
	*/

	m := map[int]string{1: "a", 3: "c", 2: "b", 4: "d", 5: "e"}
	s := make([]int, len(m))

	//对map中的key进行排序
	i := 0
	for k := range m {
		s[i] = k
		i++

	}
	sort.Ints(s)
	fmt.Println(s)

	//key和value的对调
	fmt.Println(m)
	m3 := make(map[string]int)

	for k, v := range m {
		m3[v] = k
	}
	fmt.Println(m3)
}
