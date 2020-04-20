package empty_interface

import (
	"fmt"
	"testing"
)

func typeAssertion(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Printf("value d : %d", v)
		return
	}
	if v, ok := i.(string); ok {
		fmt.Printf("value s : %s", v)
		return
	}
	fmt.Println("unKnow")
}

func TestTypeAssertion(t *testing.T) {
	typeAssertion(10)
	typeAssertion("10")
}
