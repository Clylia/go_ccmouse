package main

import "fmt"

/**
 *********************************************
 * 指针
 *********************************************
 */
// 值传递
func pass_by_value(a, b int) {
	b, a = a, b
}

func pass_by_value1(a, b int) (int, int) {
	return b, a
}

// 指针传递
func pass_by_pointer(a, b *int)  {
	*b, *a = *a, *b
}

func main() {
	// 值传递案例
	a, b := 3, 4
	pass_by_value(a, b)
	fmt.Println(a, b)

	c, d := 5, 6
	c, d = pass_by_value1(c, d)
	fmt.Println(c, d)

	// 指针传递案例
	e, f := 7, 8
	pass_by_pointer(&e, &f)
	fmt.Println(e, f)

	// 值传递返回值接收方式
}
