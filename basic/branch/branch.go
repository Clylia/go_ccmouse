package main

/**
|-------------------------------------------------------------------------------------------------
|	go 条件语句 if, switch
|-------------------------------------------------------------------------------------------------
*/

import (
	"fmt"
	"io/ioutil"
)

/**
 *********************************************
 * if 判断
 *********************************************
 */
func ifFunc() {
	const filename = "./src/learn_go_ccmouse/abc.txt"	// 还为实现读取到文件
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%s\n", contents)
	}

	// 上面案例的另一种写法
	if contents, err := ioutil.ReadFile(filename); err != nil {	// 此时的到的 contents 只能在 if 代码块里面使用
		fmt.Println(err)
	} else {
		fmt.Println("%s\n", contents)
	}
}

/**
 *********************************************
 * switch 判断
 * switch 会自动 break， 除非使用 fallthrough，这在其他语言中是没有且相反的
 *********************************************
 */
// switch 案例一，
func eval(a, b int, op string) int  {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)		// panic() 报错
	}
	return result
}

// switch 案例二，区别是 switch 关键字后面没有变量
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main()  {
	ifFunc()
	fmt.Println(eval(6, 6, "*"))
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(00),
		grade(100),
		// grade(-3),	// 将会抛出错误
	)
}