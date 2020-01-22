package main

/**
|-------------------------------------------------------------------------------------------------
|
|	panic 和 recover
|
|	panic 是强烈的错误抛出，会引起程序中断
|	recover 可以尝试把 panic 抛出的错误捞出进行分析，如果可以处理的话就进行处理，这样不致于程序中断，让程序更加健壮
|	recover 必须在 defer 的函数中使用
|
|-------------------------------------------------------------------------------------------------
*/


import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(err)
		}
	}()
	// 上面的语法是定义了一个匿名函数，转换为正常模式的话类似与 defer func funcName()，只是把 funcName 写成了一个匿名函数

	// 校验 recover 的功能
	// panic(errors.New("this is a custom error"))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic(233)
}

func main() {
	tryRecover()
}
