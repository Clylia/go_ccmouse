package main

import (
	"fmt"
	"learn_go_ccmouse/14_队列"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmtpy())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmtpy())

	q.Push("loedan")
	q.Push("ccmouse")
	fmt.Println(q)
}
