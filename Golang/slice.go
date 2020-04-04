package main

import "fmt"

func updateSlice(slice []int) {
	slice[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6]) //左闭右开 arr[2]-arr[5]
	fmt.Println("arr[:6] = ", arr[:6])   //0-6     arr[0]-arr[5]
	fmt.Println("arr[2:] = ", arr[2:])   //2-8     arr[2]-arr[7]
	fmt.Println("arr[:] = ", arr[:])     //0-8     arr[0]-arr[7]

	fmt.Println("-----------------After updateSlice s1-----------------")
	s1 := arr[2:]
	fmt.Println(s1) //[2 3 4 5 6 7]
	updateSlice(s1)
	fmt.Println(s1, arr) //[100 3 4 5 6 7] [0 1 100 3 4 5 6 7]

	fmt.Println("-----------------After updateSlice s2-----------------")
	s2 := arr[:5]
	fmt.Println(s2) //[0 1 100 3 4]
	updateSlice(s2)
	fmt.Println(s2, arr) //[100 1 100 3 4] [100 1 100 3 4 5 6 7]

	fmt.Println("-----------------Reslice s2-----------------")
	s2 = s2[:4]
	fmt.Println(s2) //[100 1 100 3]
	s2 = s2[2:]
	fmt.Println(s2) //[100 3]

	fmt.Println("-----------------Extending slice----------------")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	//fmt.Println(s1[4])  //index out of range
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d   ", s1, len(s1), cap(s1)) //s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))  //s2=[5 6], len(s2)=2, cap(s2)=3

	fmt.Println("------------------slice append-----------------")
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	//s4 and s5 no longer view arr.
	fmt.Println("before modify s4, arr = ", arr) //0 1 2 3 4 5 6 10
	s4[0] = 100
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d   ", s3, len(s3), cap(s3)) // s3=[5 6 10], len(s3)=3, cap(s3)=3
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d   ", s4, len(s4), cap(s4)) // s4=[100 6 10 11], len(s4)=4, cap(s4)=6
	fmt.Printf("s5=%v, len(s5)=%d, cap(s5)=%d\n", s5, len(s5), cap(s5))  // s5=[100 6 10 11 12], len(s5)=5, cap(s5)=6
	fmt.Println("behind modify s4，arr = ", arr)                          //0 1 2 3 4 5 6 10
}
