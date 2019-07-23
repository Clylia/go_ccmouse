package main



import "fmt"

var lastOccurred = make([]int, 0xffff)	// 对上面代码的优化，用空间换时间，提高运行速度

// 算法：得出一个字符串中连续不重复的字串的最大数
func lengthOfNonRepeatingSubStr(s string) int {
	// lastOccurred := make(map[rune]int)
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main()  {
	// 例题，得出一个字符串中连续不重复的字串的最大数
	fmt.Println(lengthOfNonRepeatingSubStr("lalala"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcddcac"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
