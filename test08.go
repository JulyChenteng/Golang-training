package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}

	for _, f := range fs {
		f()
	}
}

/*
	closure i =  4
	closure i =  4
	closure i =  4
	closure i =  4
	defer_closure i =  4
	defer i =  3
	defer_closure i =  4
	defer i =  2
	defer_closure i =  4
	defer i =  1
	defer_closure i =  4
	defer i =  0
**/
