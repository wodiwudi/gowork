package main

import (
	"fmt"
	"test/retriever/mock"
	real2 "test/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, data map[string]string) string
}

type RetrieverPoster interface { //接口的组合
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func downloadGet(r Retriever) string {
	return r.Get(url)
}

func downloadPost(p Poster) string {
	return p.Post(url, map[string]string{
		"name": "zxw",
	})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"context": "another imooc faked",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	//switch判断类型
	fmt.Println(r)
	switch v := r.(type) {
	case *mock.MyRetriever:
		fmt.Println("mockRetriever", v.Context)
	case *real2.Myretriever:
		fmt.Println("realRetriever", v.UserAgent)
	}

	//Type assertion（断言判断类型）
	//v := r.(*real2.Myretriever)
	//fmt.Println(v.UserAgent)
}

func main() {
	var r Retriever
	r = &real2.Myretriever{
		UserAgent: "apache",
		TimeOut:   time.Minute,
	}
	r = &mock.MyRetriever{Context: "gg"}
	inspect(r)
	fmt.Println("***********************")
	var rs RetrieverPoster
	rs = &mock.MyRetriever{Context: "initial value"}
	res := session(rs)
	fmt.Println(res)
}
