package main

import (
	"fmt"
	"learn_go_ccmouse/15_接口/mock"
	"learn_go_ccmouse/15_接口/really"
	"time"
)

/**
|-------------------------------------------------------------------------------------------------
|	go 接口 interface
|
|	    go 语言中接口（interface）只要在自己内部定义好需要的方法名以及方法需要的参数和返回值类型
|   即表示定义好了接口，
|		接口的实现不需要像 php 一样用关键字 implements 来实现，而是当一个结构体将一个接口里面所
|   有的方法都实现的时候，就表示这个结构体继承了这个接口
|-------------------------------------------------------------------------------------------------
*/

/**
 * Retriever 接口
 */
// 定义一个接口
type Retriever interface {
	// 在接口中的方法前面不需要 func 关键字
	Get(url string) string
}

const url = "http://www.imooc.com"

// 调用 Retriever 接口中的方法
func download(r Retriever) string {
	return r.Get(url)
}


/**
 * Poster 接口
 */
type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url, map[string]string {
		"name": "loedan",
		"course": "golang",
	})
}


/**
 * 接口组合，为了实现多个接口的功能
 */
type RetrieverPoster interface {
	// 把要实现的接口放进来，就可以将接口组合在一起，但是多个接口不可以有重名的方法，
	Retriever
	Poster

	// 组合接口中也可以声明方法
	//Connect(host string) string
}

func session(s RetrieverPoster) string {
	// 组合接口拿进来后可以直接使用各个接口中的方法
	// s.Get()
	s.Post(url, map[string]string {
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)

	// really.Retriever 这个结构体里面有一个 Get() 方法有一个指针接收者，所以这里要加上 & 来取地址
	r = &really.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion 类型断言
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("Not a mock retriever")
	}


	// fmt.Println(download(r))		// 执行方法
	fmt.Println("Try s session")
	// session() 方法里面需要的是一个 RetrieverPoster 类型的参数，
	// 既要求实现 Retriver 接口里面的 Get() 方法，又要求实现 Poster 接口里面的 Post() 方法
	// 而此时的 retriever 变量是 mock.Retriever{}，这个结构体刚好将这两个方法都实现了，
	// 所以可以直接传进 session() 方法中使用
	fmt.Println(session(&retriever))		// 执行方法
}

// 判断类型
func inspect(r Retriever)  {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *really.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
