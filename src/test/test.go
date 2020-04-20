package main

import (
	"fmt"
)

type RedisKey string

func (k RedisKey) String() string {
	return string(k)
}

//原生格式化
func (k RedisKey) Fmt(args ...interface{}) string {
	return fmt.Sprintf(k.String(), args...)
}
func main() {
	var aa RedisKey
	aa = "%s123%d"
	bb := aa.Fmt("g", 2)
	fmt.Println(bb)

}
