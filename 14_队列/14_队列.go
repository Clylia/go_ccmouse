package queue


/**
|-------------------------------------------------------------------------------------------------
|
|	设置为 interface{}，在go语言中表示支持任何类型
|
|-------------------------------------------------------------------------------------------------
*/
// 一个接口（interface）类型的切片，表示这个切片里面可以放任何类型的内容
type Queue []interface{}

// 塞入	Pushes the element into the queue
// 	测试文档功能 e.g. q.Push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// 弹出	Pops element from head
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]

	// 由于 Queue 这个类型里面内容是 interface 类型的内容，可以是任何类型的元素
	// 而当前函数返回值需要的是 int 类型的值
	// 所以使用 .(int) 将返回值强制转为 int 类型
	return head.(int)
}

// 判断是否为空	Returns if the queue is empty or not
func (q *Queue) IsEmtpy() bool {
	return len(*q) == 0
}
