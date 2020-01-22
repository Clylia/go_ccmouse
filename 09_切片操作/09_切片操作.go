package main

import "fmt"

/**
|-------------------------------------------------------------------------------------------------
|	go 对 slice 的添加，删除，修改等操作
|
|	make([]int, 10, 32)	// 创建一个切片，10 为切片的 len， 32为切片的 cap，如果参数 3 不传，则默认与 len 相同
|
|-------------------------------------------------------------------------------------------------
*/
// 打印 slice 信息
func printSlice(s []int) {
	fmt.Printf("value=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

// 创建 slice
func createSlice()  {
	fmt.Println("Create Slice")
	// 普通方式创建 slice
	var s []int	// Zero value for slice is nil, nil 为没有，等同与 PHP 或 javascipt 的 null
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i * 2 + 1)
	}
	fmt.Println(s)

	// 简易方式创建 slice
	s1 := []int{2, 4, 6, 8}	// 创建一个数组，同时创建一个切片
	printSlice(s1)
	s2 := make([]int, 16)	// 16 为切片的 len
	printSlice(s2)
	s3 := make([]int, 10, 32)	// 10 为切片的 len，32 为切片的 cap 值
	printSlice(s3)
}

// 拷贝 slice
func copySlice()  {
	fmt.Println("Copy Slice")
	s := []int{1, 2, 3, 4, 5}
	s1 := make([]int, 10, 32)
	copy(s, s1)		// 把参数二的内容放进参数一的内容开头，并覆盖对应长度的内容
	printSlice(s)
}

// 删除 slice 的元素
func deleteSlice()  {
	fmt.Println("删除头部元素")
	s := []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	front := s[0]
	s = s[1:]
	fmt.Println(front)
	printSlice(s)

	fmt.Println("删除中间元素")
	s1 := []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	// go 中没有内置的删除切片元素的方法，但是可以通过 append 的方法来实现
	s1 = append(s1[:3], s1[4:]...)	// append() 方法后面是一个可变的参数，此时在第二个参数后面加上 ... 即可
	printSlice(s1)

	fmt.Println("删除尾部元素")
	s2 := []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(tail)
	printSlice(s2)
}

func main() {
	createSlice()	// 创建 slice
	copySlice()		// copy slice
	deleteSlice()	// 删除 slice 中的元素
}
