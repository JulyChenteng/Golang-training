package main

import "fmt"

func PrintArray(arr [5]int) {
	arr[0] = 1000
	for _, v := range arr {
		fmt.Println(v)
	}
}

func PrintArray2(arr *[5]int) {
	arr[0] = 1000
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//for i:=0; i<len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}

	//for _, v := range arr3 {
	//	fmt.Println(v)
	//}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	PrintArray(arr3)  //1000 4 6 8 10 函数中处理的是数组的拷贝
	fmt.Println(arr3) //2 4 6 8 10

	PrintArray2(&arr3) //1000 4 6 8 10 通过数组指针来修改数组内容
	fmt.Println(arr3)  //1000 4 6 8 10
}
