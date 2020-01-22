package main

import "fmt"

/**
|-------------------------------------------------------------------------------------------------
|	go slice 切片
|
|	半开半闭规则，左边下标包含在结果内，右边下标不包含在结果内
|	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
|	s := arr[2:6]	// s的值为[2 3 4 5]，
|	s := arr[:6]	// s的值为[0 1 2 3 4 5]，开始下标省略后，表示从 0 开始
|	s := arr[2:]	// s的值为[2 3 4 5 6 7]， 结束下标省略后，表示截取到最后一个
|	s := arr[:]		// s的值为[0 1 2 3 4 5 6 7]， 开始和结束下标都省略的话，表示从数组开始到结束都截取
|
|	slice 是对原本 arr 的一个 view，改变 slice 的内容，原 arr 对应的内容也会改变
|	arr := [...]int{1, 2, 3, 4, 5}
|	s := arr[1:3]	// [2 3]
| 	len(s)	// 2 切片的len
| 	cap(s)	// 3 切片开始位置在数组中的位置到数组的结尾的len
|	s = append(s, 1)	// 像切片中添加元素，此时的 arr = [1, 2, 3, 1, 5]，添加元素时，如果超越 cap，系统会重新分配更大的底层数组，
|					// 原来的数组如果有地方使用，则arr还存在，如果没有地方使用，go 会使用垃圾回收机制将原 arr 垃圾回收掉
|					// 由于 slice 是引用传递，必须接收 append 的返回值
|-------------------------------------------------------------------------------------------------
*/
// 基础操作切片及理解
func basic() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])
}

// 更新切片内容，切片是引用传递，更改了切片内容，原数组内容也将被改变
func updateSlice(s []int) {
	s[0] = 100
	fmt.Println(s)
}

// 对切片执行切片动作
func reslice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Reslice:")
	s := arr[:]
	s = s[:5]	// 对切片执行切片，为 reslice 动作
	fmt.Println(s)
	s = s[2:]
	fmt.Println(s)
}

// 切片可以向后扩展
func extendSlice() {
	fmt.Println("Extend Slice:")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	s1 := arr[2:6]	// [2 3 4 5]
	s2 := s1[3:5]	// [5 6]，可向后扩展，但是 s[i] 不可以超越 len(s),向后扩展不可以超越底层数组 cap(s)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))	// s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))	// s2=[5 6], len(s2)=2, cap(s2)=3

	// 像 slice 中添加元素
	fmt.Println("像 slice 中添加元素")
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5=", s3, s4, s5)
	// s3	[5 6 10]
	// s4	[5 6 10 11]
	// s5	[5 6 10 11 12]
	fmt.Println(arr)
	// [0 1 2 3 4 5 6 10]，这个 arr len只有 7，所以 s4 和 s5 添加的元素不会放置在 arr 中，
	// 系统会把这个 arr 拷贝一份，此时，s4 和 s5 添加的元素会放置在拷贝出来的新 arr 中
}

func main()  {
	// 基础操作
	basic()

	// 验证切片为 引用传递
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:]
	s2 := arr[:]
	// 先执行 s1
	updateSlice(s1)
	fmt.Println("After slice s1")
	fmt.Println(arr)
	// 再执行 s2
	updateSlice(s2)
	fmt.Println("After slice s2")
	fmt.Println(arr)

	reslice()	// 对切片执行切片动作
	extendSlice()
}
