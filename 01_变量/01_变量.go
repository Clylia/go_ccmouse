package main

/**
|-------------------------------------------------------------------------------------------------
|	go 语言基础知识
|-------------------------------------------------------------------------------------------------
*/

import (
	"fmt"
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

func main() {
	// 定义变量的案例
	variableZeroValue()
	variableInitialValue()
	varibaleTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc, dd, ee, ff)
}
