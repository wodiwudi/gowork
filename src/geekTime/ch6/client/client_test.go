package client

import (
	"geekTime/ch6/test_package"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(test_package.TestFun())
}
