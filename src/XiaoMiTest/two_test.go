package XiaoMiTest

import (
	"fmt"
	"testing"
)

//for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址
func TestTwo(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	//错误写法
	/*for key, val := range slice {
		m[key] = &val
	}*/
	for key, val := range slice {
		value := val
		m[key] = &value
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
