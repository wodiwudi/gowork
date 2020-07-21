package testing

import (
	"fmt"
	"os"
	"testing"
	"time"
)

//使用parallel可以有效利用多核并行优势
func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	if os.Args[len(os.Args)-1] == "b" {
		t.Parallel()
	}
	time.Sleep(time.Second * 2)
}

//go test -v   4s
//go test -v -args "b"  2s

func add(x, y int) int {
	return x + y
}

//output注释该函数就会执行，不过不能用print于println用fmt包的内容
func ExampleAdd() {
	fmt.Println(add(1, 2))
	fmt.Println(add(3, 4))

	//output:
	//3
	//7
}
