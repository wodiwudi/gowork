package engine

//解析结果类型结构体
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//请求结构体
type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

//不做任何事情的parser
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
