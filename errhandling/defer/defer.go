package main

/**
|-------------------------------------------------------------------------------------------------
|	资源处理 和 错误处理
|
|
|	资源处理
|		defer 关键字，带有 defer 的语句的输出会在函数结束时输出，就算程序报错了，报错前的 defer 输出也一会会输出
|		函数结束前若有多个 defer 语句，遵循 先进后出 的原则，类似于栈
|
|	错误处理
|		if err != nil {
|		// 方式一：直接抛出
|		// panic(err)
|
|		// 方式二：抛出程序获取的错误信息
|		// fmt.Println("Error:", err.Error())
|
|		// 方式三：获取错误的详细信息，让代码更强壮
|		if pathError, ok := err.(*os.PathError); !ok {
|			panic(err)
|		} else {
|			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
|		}
|		return
|	}
|
|-------------------------------------------------------------------------------------------------
*/

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

// defer 试用
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic(32)
	return
	fmt.Println(4)
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// 自定义的一个错误信息
	// err = errors.New("this is a custom error")
	if err != nil {
		// 方式一：直接抛出
		// panic(err)

		// 方式二：抛出程序获取的错误信息
		// fmt.Println("Error:", err.Error())

		// 方式三：获取错误的详细信息，让代码更强壮
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()	// 保证关闭这个文件资源

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++{
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
