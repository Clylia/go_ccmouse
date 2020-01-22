package _5_函数

/**
|-------------------------------------------------------------------------------------------------
|	go 函数与指针
|-------------------------------------------------------------------------------------------------
*/

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
 *********************************************
 * 函数的定义与不同方式
 *********************************************
 */
// 简易计算器
func calculator(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div1(a, b)	// _ 表示函数的返回值被放弃不接收
		return q
	default:
		panic("unsupported operator:" + op)		// panic() 报错
	}
}

// 带 error 返回的简易计算器
func calculator1(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div1(a, b)	// _ 表示函数的返回值被放弃不接收
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
}

// 多个返回值的函数
func div(a, b int) (int, int) {
	return a / b, a % b
}

// 多个返回值，且给返回值命名的函数
func div1(a, b int) (q, r int) {
	return a / b, a % b
}

// 多个返回值，且给返回值命名的函数
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return	// 这种形式可以返回 q 和 r，但是当函数体里面内容多的时候不便与阅读，所以不建议这样使用
}

// 函数式编程，函数的参数为函数，类似 php 中的闭包函数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()	// 利用反射得到指针
	opName := runtime.FuncForPC(p).Name()	// 从内存中得到运行的函数名
	fmt.Printf("Call function %s with args " + "(%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 重写 math 中的 Pow 方法
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表的函数
func sumArgs(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func main() {
	fmt.Println(calculator(1, 3, "-"))
	if result, err := calculator1(1, 2, "a"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(div(5, 3))
	q, r := div1(13, 3)	// 接收返回的内容
	fmt.Println(q, r)
	fmt.Println(div2(13, 3))

	fmt.Println(apply(pow, 3, 4))
	// 将闭包直接写成一个匿名函数
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sumArgs(1, 2, 3, 4, 5))
}
