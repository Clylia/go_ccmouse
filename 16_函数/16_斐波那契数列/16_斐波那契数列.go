package fib

// 斐波那契数列 实现
// 1 1 2 3 5 8 13 21...,后一个数是前两个数的和
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}
