package main

/**
|-------------------------------------------------------------------------------------------------
|	go 数组
|
|	[10] int 和 [20] int 是不同类型
|	调用 func f(arr [10] int) 会在内存中拷贝数组，而不是对原数组做处理
|-------------------------------------------------------------------------------------------------
*/

import "fmt"

// 创建数组
func createArrays() {
	var arr1 [5]int		// 5 表示数组的长度，未设默认值时，int 类型的默认值统一为 0
	arr2 := [3]int{1, 3, 5}	// 简易方式创建数组，必须赋初始值
	arr3 := [...]int{2, 4, 6, 8, 10}	// ... 表示让编译器自动计算数组长度
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// 二维数组
	var grid [4][5] int	// 定义一个四行五列的二维数组
	fmt.Println(grid)
}

// 数组的遍历，range 实现遍历
func traverseArray() {
	arr := [...] int {1, 2, 3, 4, 5}
	// 常规方式
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, " => ", arr[i])	// i 为数组下标
	}

	// range 方式，只传递下标
	for i := range arr {
		fmt.Println(i, " => ", arr[i])
	}

	// 即传递下标，也传递值
	for i, v := range arr {
		fmt.Println(i, " => ", v)	// v 为下标对应的值
	}

	// 放弃下标，传递值
	for _, v := range arr {
		fmt.Println("value => ", v)	// v 为下标对应的值
	}
}

// 值传递打印数组内容
func printArrayByValue(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, " => ", v)
	}
}
// 引用传递打印数组内容
func printArrayByPointer(arr *[6]int)  {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, " => ", v)
	}
}

func main()  {
	createArrays()		// 创建数组
	traverseArray()		// 遍历数组

	// 默认的是值类型传递，会在内存中拷贝一份，不会影响原数组内容
	arr := [5]int{0, 2, 4, 6, 8}
	printArrayByValue(arr)
	fmt.Println(arr)

	// 把地址指针传递过去
	arr1 := [6]int{0, 2, 4, 6, 8, 10}
	printArrayByPointer(&arr1)
	fmt.Println(arr1)
}
