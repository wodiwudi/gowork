package XiaoMiTest

import (
	"fmt"
	"testing"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func TestTwelve(t *testing.T) {
	tt := Teacher{}
	//通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。此外，外部类型还可以定义自己的属性和方法，
	//甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。这个例子中的 ShowB() 就是同名方法
	tt.ShowB() //使用的是外部的showB，people的showB被屏蔽了
	//可以
	tt.People.ShowB() //可以详细调用来使用内部的同名函数

	tt.ShowA() //showA showB Teacher没有showA方法，所以会调用内部的showA

	str := "hello"
	//str[0] = 'x' 字符串是只读的
	bytes := []byte(str) //先转化为byte或rune数组再改再转换为string
	bytes[0] = 'x'
	str = string(bytes)
	fmt.Println(str)
}
