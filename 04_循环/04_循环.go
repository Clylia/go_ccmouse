package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/**
|-------------------------------------------------------------------------------------------------
|	go 循环
|-------------------------------------------------------------------------------------------------
*/
// 十进制转换成二进制函数
func convertToBin(n int) string {
	result := ""
	// 这个 for 循环省略了起始条件
	for ; n > 0; n /= 2 {	// n /= 2，整形与整形执行数学运算得到的结果也是整形，如果有小数点，则会舍弃
		//fmt.Println(n)
		lsb := n % 2
		result = strconv.Itoa(lsb) + result		// strconv.Itoa() int 型转换成 string 类型
	}
	return result
}

// 使用 for 循环逐行读取文件内容
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)	// 程序中断，抛出错误
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	// 这个 for 循环省略了 初始条件、递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 建立一个死循环
func forever() {
	// 这个循环任何条件都没有，是一个死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),	// 101
		convertToBin(13),	// 1101
		convertToBin(123456),	// 11110001001000000
		convertToBin(0),
	)

	printFile("./04_循环/abc.txt")

	s := `abc"d"
	kkk
	loedan
	p`
	printFileContents(strings.NewReader(s))

	//forever()
}