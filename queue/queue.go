package queue


/**
|-------------------------------------------------------------------------------------------------
|
|	设置为 interface{}，在go语言中表示支持任何类型
|
|-------------------------------------------------------------------------------------------------
*/
// An FIFO queue
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
	return head.(int)	// 将返回值强制转为 int 类型
}

// 判断是否为空	Returns if the queue is empty or not
func (q *Queue) IsEmtpy() bool {
	return len(*q) == 0
}
