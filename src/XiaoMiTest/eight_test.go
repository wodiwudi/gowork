package XiaoMiTest

import (
	"fmt"
	"testing"
)

func hello() []string {
	return nil
}
func TestEight(t *testing.T) {
	h := hello    //相当于把函数赋值给h,h的类型为函数 ，如果是h()则为nil
	h1 := hello() //相当于把hello()的返回值nil赋值给h1
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	if h1 == nil {
		fmt.Println("h1 nil")
	} else {
		fmt.Println(" h1 not nil")
	}
	fmt.Printf("%T", h) //func() []string
}

func TestEight2(t *testing.T) {
	i := -5
	j := +5
	fmt.Printf("%+d %+d", i, j) //+表示输出数值的符号
	fmt.Println()
	fmt.Printf("%d %d", i, j)
}
