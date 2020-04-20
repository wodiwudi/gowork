package engine

import (
	"log"
	"test/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	//循环查找
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)

		bytes, e := fetcher.Fetch(r.Url)
		//如果有错误则跳过
		if e != nil {
			log.Printf("Fetcher:error fetchin url ：%s,err ：%v", r.Url, e)
			continue
		}
		//把其他的request放进requests数组
		parseResult := r.ParseFunc(bytes)
		requests = append(requests, parseResult.Requests...)
		//打印items
		for _, item := range parseResult.Items {
			log.Printf("item:%v", item)
		}
	}
}
