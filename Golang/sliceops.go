package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Println("s=", s, "len=", len(s), "cap=", cap(s))
}

func main() {
	fmt.Println("---------------------Creating slice-------------------------")
	var s []int   // Zero value for slice is nil
	printSlice(s) //s= [] len= 0 cap= 0

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	printSlice(s) //s= [1 3 5 ... 199] len= 100 cap= 128

	fmt.Println()
	s1 := []int{2, 4, 6, 8}
	printSlice(s1) //s= [2 4 6 8] len= 4 cap= 4

	s2 := make([]int, 16) //s= [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] len= 16 cap= 16
	printSlice(s2)

	s3 := make([]int, 16, 32) //s= [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] len= 16 cap= 32
	printSlice(s3)

	fmt.Println("------------------------------Copying slice----------------------")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("-------------------Deleting elements from slice------------------")
	s2 = append(s2[:3], s2[4:]...) //删除s2[3]，append的第二个参数为可变长参数，以s2[4:]...表示用该slice的所有元素做可变长参数
	printSlice(s2)

	fmt.Println("-------------------Popping from front--------------------")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("-------------------Popping fron back---------------------")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}
