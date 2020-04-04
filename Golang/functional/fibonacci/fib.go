package fibonacci

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type FibFunc func() int

func Fib() func() int {
	a, b := -1, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

//实现io.Reader接口
func (f FibFunc) Read(p []byte) (n int, err error) {
	next := f()
	if next > 10000 {
		return 0, io.EOF
	}
	fmt.Println(next)
	s := fmt.Sprintf("%d\n", next)

	//TODO : incorrect if p is too small!!!
	return strings.NewReader(s).Read(p)
}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*fibonacci另一种实现*/
type iFib func() (int, iFib)

func Fib2(f1, f2 int) iFib {
	return func() (int, iFib) {
		return f1 + f2, Fib2(f2, f1+f2)
	}
}
