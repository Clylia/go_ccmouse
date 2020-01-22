package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

/**
 * resp, err := http.Get("http://www.baidu.com")  // 直接请求一个网址并获得返回内容
 * request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)	// 创建一个请求
 * request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
 * resp, err := http.DefaultClient.Do(request)  // 默认 client 执行请求
 */

func main() {
	//resp, err := http.Get("http://www.baidu.com")  // 直接请求一个网址并获得返回内容
	
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)	// 创建一个请求
	// 添加头部信息
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//resp, err := http.DefaultClient.Do(request)  // 默认 client 执行请求
	
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
