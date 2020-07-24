package XiaoMiTest

import (
	"fmt"
	"testing"
)

type MyInt1 int   //相当于创建新类型
type MyInt2 = int //相当于创建别名

func TestSix(t *testing.T) {
	var i int = 0
	//var i1 MyInt1 = i  //强类型语言下不通过
	var i2 MyInt2 = i //赋值给别名可以通过
	fmt.Println(i2)
}
