package main

import "fmt"

/******************************************************
 * 计算一个字符串中最长不含邮重复字符字串的长度
 *****************************************************/

var lastPo = make([]int, 0xffff)  // 0xffff 16进制数

func getNonRepeatStrByByte(s string) int {
	// 字符串的每一个字节的类型为 byte，golang 中没有字符（char）这个类型
	// 这个 lastPo 预设要存的数据为 {'a': 0, 'b': 1, 'c': 3...}，所以 map 的 key 为 byte
	//lastPo := make(map[byte]int)
	for i := range lastPo {
		lastPo[i] = -1
	}

	start := 0
	maxLength := 0

	// []byte(s) 是将字符串 s 转换成 byte 类型的数组或切片
	for i, ch := range []byte(s) {
		if lastI := lastPo[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastPo[ch] = i
	}

	return maxLength
}

func getNonRepeatStrByRune(s string) int {
	// 这个 lastPo 预设要存的数据为 {'a': 0, 'b': 1, 'c': 3...}，所以 map 的 key 为 byte
	//lastPo := make(map[rune]int)	// 将 map 的 key 定义为 rune 类型
	for i := range lastPo {
		lastPo[i] = -1
	}

	start := 0
	maxLength := 0

	// []rune(s) 是将字符串 s 转换成 rune 类型的数组或切片
	for i, ch := range []rune(s) {
		if lastI := lastPo[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastPo[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(getNonRepeatStrByRune("lalala"))
	fmt.Println(getNonRepeatStrByRune("bbb"))
	fmt.Println(getNonRepeatStrByRune("abcddcac"))
	fmt.Println(getNonRepeatStrByRune(""))
	fmt.Println(getNonRepeatStrByRune("b"))
	fmt.Println(getNonRepeatStrByRune("abcdef"))
	fmt.Println(getNonRepeatStrByRune("这里是慕课网"))
	fmt.Println(getNonRepeatStrByRune("一二三二一"))
	fmt.Println(getNonRepeatStrByRune("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
