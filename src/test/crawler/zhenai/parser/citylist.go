package parser

import (
	"regexp"
	"test/crawler/engine"
)

const cityListRegexp = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9:/.]+)"[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRegexp)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))         //Item里方城市列表名
		result.Requests = append(result.Requests, engine.Request{ //Requests放下一个请求
			Url:       string(m[1]),
			ParseFunc: engine.NilParser,
		})
	}
	return result
}
