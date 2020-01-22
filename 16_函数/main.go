package main

import (
	"bufio"
	"fmt"
	"io"
	fib "learn_go_ccmouse/16_函数/16_斐波那契数列"
	"strings"
)

/**
 * 一个函数也可以作为一个类型
 */
type intGen func() int

// 函数的类型也可以实现接口，也可以作为接收者
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	// 这个 for 循环省略了 初始条件、递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()

	printFileContents(f)
}
