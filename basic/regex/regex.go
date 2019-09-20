package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is loedan@163.com@abc.com
email is abc@aaa.com
email2 is kkk@imooc.com
`

// 正则表达式
func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
