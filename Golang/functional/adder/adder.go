package adder

func Adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func Adder2(base int) iAdder {
	return func(num int) (int, iAdder) {
		return base + num, Adder2(base + num)
	}
}
