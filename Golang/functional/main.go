package main

import (
	"./fibonacci"
	"fmt"
)

func main() {
	//a := adder.Adder()
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	//}

	//a1 := adder.Adder2(0)
	//for i := 0; i < 10; i++ {
	//	var s int
	//	s, a1 = a1(i)
	//	fmt.Println(i, "  " , s)
	//}

	var fib fibonacci.FibFunc
	fib = fibonacci.Fib()
	fmt.Println(fib(), fib(), fib(), fib())

	fib2 := fibonacci.Fib2(-1, 1)
	for i := 0; i < 8; i++ {
		var res int

		res, fib2 = fib2()
		fmt.Print(res, " ")
	}
	fmt.Println()

	fibonacci.PrintFileContents(fib)
}
