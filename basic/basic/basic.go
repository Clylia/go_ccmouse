package main

/**
|-------------------------------------------------------------------------------------------------
|	go 语言基础知识
|-------------------------------------------------------------------------------------------------
*/

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
 *********************************************
 * 1、 定义变量
 *********************************************
 */
// 函数外部的变量在包内的任何函数内都可以使用，
// 函数外部变量的在定义时必须严格按照 var name = xxx 格式，不允许使用简写方式 := 来定义，简写方式只允许在函数内部使用
// 函数外部变量多个时可直接用 （）阔起来，让代码看起来更简洁
var aa = 3
var bb = true
var cc = "kkk"
var (
	dd = 2
	ee = false
	ff = "this is a string"
)

// 初始化为 0 或 空字符串的变量定义
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n",a, s)
}

// 有初始值的变量定义
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 自动识别变量类型的变量定义方式
func varibaleTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 变量的简写方式
func variableShorter() {
	// := 的简写方式只能在函数内部使用
	// 在函数内部定义变量尽量使用简写方式，可以简化代码
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}


/**
 *********************************************
 * 2、 数据类型
 * bool, string
 * (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr（指针）  // 带 u 表示无符号整数， 不带 u 表示有符号整数
 * byte， rune
 * float32, float64, complex64, complex128	// complex：复数
 * 类型转换是强制的
 *********************************************
 */
// 欧拉公式代码实现
func euler()  {
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

// 强制类型转换
func triangle()  {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	// math.Sqrt 需要 float64 的参数，并返回 float64 的参数，而 c 定义的是 int 类型，所以接收返回值时需要强制转换
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}


/**
 *********************************************
 * 3、 常量与枚举
 * const = "abc.txt"
 * const 数值可作为各种类型使用
 * go 中定义常量建议不要像其他语言一样全部使用大写，在 go 中，大小写是有可见性含义的
 * iota 表示这组 const 是自增值的，后面的就不需要定义赋值
 *********************************************
 */
const filename = "abc.txt"	// 方法外定义常量
// 定义常量
func consts()  {
	const a, b = 3, 4	// 方法内定义常量
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(c)

	// 常量也可以括号里面定义多个，简化代码
	const (
		name = "loedan"
		age = 18
	)
	fmt.Println(name, age)
}

// 枚举类型
func enums() {
	// go 语言没有单独的枚举类型，通过 const 定义一组枚举类型
	const (
		zero = 0
		one = 1
		two = 2
		three = 3
	)
	fmt.Println(zero, one, two, three)

	const (
		php = 0
		cpp = iota	// iota 表示这组 const 是自增值的，后面的就不需要定义赋值
		_
		python
		golang
		javascript
	)
	fmt.Println(php, cpp, javascript, python, golang)

	// iota 参与运算
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	// 定义变量的案例
	variableZeroValue()
	variableInitialValue()
	varibaleTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc, dd, ee, ff)

	// 数据类型的案例
	euler()
	triangle()

	// 常量与枚举
	consts()
	enums()
}
