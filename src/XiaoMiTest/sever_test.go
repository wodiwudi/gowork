package XiaoMiTest

import (
	"fmt"
	"testing"
)

//iota必须出现在const及const的代码块中，从零开始计数，可以看成是从零开始的const的索引
const (
	x = iota
	_
	y
	z = "zz"
	k //注意，当const中不赋值时为上一行相同的值，如果上一行是iota则本行为加一
	p = iota
)

//每次 const 出现时，都会让 iota 初始化为0.
const (
	b = iota //b=0
	c        //c=1
)

func TestSeven(t *testing.T) {
	fmt.Println(x, y, z, k, p, b, c)
	//0 2 zz zz 5 0 1
}
