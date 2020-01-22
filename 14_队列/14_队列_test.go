package queue

/**
|-------------------------------------------------------------------------------------------------
|
|	示例的方法全部以 Example 开头
|
|-------------------------------------------------------------------------------------------------
*/

import "fmt"

// 示例
func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmtpy())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmtpy())

	// Output：
	// 1
	// 2
	// false
	// 3
	// true
}
