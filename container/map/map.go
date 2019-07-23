package main

/**
|-------------------------------------------------------------------------------------------------
|	go 对 hash map 的定义与使用
|
|
|	map 的结构：map[key]value，其中 key 和 value 都是数据类型
|	m1 := map[string]string {	// string 类型的 key ，string 类型的 value
|		"name": "loedan",
|		"course": "golang",
|		"site": "imooc",
|		"quality": "notbad",
|	}
|
|	map 中的内容是无序的，每次取同一个 map 得到的值顺序可能不一致
|
|-------------------------------------------------------------------------------------------------
*/

import "fmt"

// 创建一个 map，意思为 hash map
func createMap()  {
	// 普通定义方式，定义一个 map 并初始化赋值
	m1 := map[string]string {
		"name": "loedan",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	fmt.Println(m1)

	// 用 make 创建一个 map
	m2 := make(map[string]int)
	fmt.Println(m2)

	// 初始化一个空 map
	var m3 map[string]int
	fmt.Println(m3)
}

// 遍历 map
func traverserMap() {
	m := map[string]string {
		"name": "loedan",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// 从 map 中获取元素
func getMapElement()  {
	fmt.Println("get element form map")
	m := map[string]string {
		"name": "loedan",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	// fmt.Println(abc)	// 会的到一个 zero value，go 在获取 map 中不存在的元素时，不会报错
	// 校验获取 map 中的元素是否存在
	if abc, ok := m["abc"]; ok {
		fmt.Println(abc)
	} else {
		fmt.Println("key does not exist")
	}
}

// 删除 map 中的元素
func deleteMapElement() {
	fmt.Println("delete element form map")
	m := map[string]string {
		"name": "loedan",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	name, ok := m["quality"]
	fmt.Println(name, ok)
	delete(m, "quality")
	fmt.Println(m)
	fmt.Println(name, ok)
}

func main() {
	createMap()			// 创建 map
	traverserMap()		// 遍历 map
	getMapElement()		// 从 map 中获取元素
	deleteMapElement()	// 删除 map 中的元素
}
