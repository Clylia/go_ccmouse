package main

/**
|-------------------------------------------------------------------------------------------------
|	go 字符串类型 string
|
|	rune 的讲解
|		rune 相当与 go 的 char
|		使用 range 遍历 pos, rune 对
|		使用 utf8.RuneCountInString() 获得字符数量
|		使用 len() 获得字节长度
|		使用 []byte 获得字节
|		使用 []rune 获得 rune
|
|	字符串类型的函数包为 strings，
|
|-------------------------------------------------------------------------------------------------
*/

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	fmt.Printf("%s\n", []byte(s))
	fmt.Printf("%X\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	// 循环打印的结果
	// 59 65 73 E6 88 91 E7 88 B1 E6 85 95 E8 AF BE E7 BD 91 21
	fmt.Println()

	for i, ch := range s {	// ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	// 循环打印结果
	// (0 59) (1 65) (2 73) (3 6211) (6 7231) (9 6155) (12 8BFE) (15 7F51) (18 21)
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))	// 获取一个字符串的 rune 类型的长度

	// 解读 rune
	bytes := []byte(s)
	fmt.Println(bytes)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	// 循环打印的结果
	// Y e s 我 爱 慕 课 网 !
	fmt.Println()

	// 将字符串转换成 rune 类型，就可以通过下标获取对应的字符
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	// 循环打印的结果
	// (0 Y) (1 e) (2 s) (3 我) (4 爱) (5 慕) (6 课) (7 网) (8 !)
	fmt.Println()
}
