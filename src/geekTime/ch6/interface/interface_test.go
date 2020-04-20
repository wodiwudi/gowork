package _interface

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteCode() string
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (gp *GoProgrammer) WriteCode() string {
	return "go语言里面的实现接口"
}

func (jp *JavaProgrammer) WriteCode() string {
	return "JAVA语言里面的实现接口"
}

func TestGoProgrammer(t *testing.T) {
	gp := new(GoProgrammer)
	jp := &JavaProgrammer{}
	fmt.Println(gp.WriteCode())
	fmt.Println(jp.WriteCode())
}
