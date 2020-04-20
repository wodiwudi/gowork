package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	//改进版：加上Header头内容 获得手机版的imooc内容
	request, e := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil) //新建一个request请求 get无body
	if e != nil {
		panic(e)
	}
	//添加request头内容
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) "+
		"AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	response, err := http.DefaultClient.Do(request) //用默认客户端函数接收request
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() //关闭响应体

	s, err := httputil.DumpResponse(response, true) //将响应输出
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
