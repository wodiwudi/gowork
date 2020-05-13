package dairyTest

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestOne(t *testing.T) {
	testStr := "张校..炜"

	for i := 0; i < len(testStr); i++ {
		fmt.Printf("%d,%c\n", i, testStr[i])
	}
	for k, v := range testStr {
		fmt.Printf("%d,%c\n", k, v)
	}
	fmt.Println(utf8.RuneCountInString(testStr))
	stringSum2()
}

func stringSum() string {
	res := ""
	for i := 0; i < 1000; i++ {
		res += "a"
	}
	return res
}

func stringSum2() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}
func stringSum3() string {
	var b bytes.Buffer
	b.Grow(1000)
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}
func BenchmarkTestStringSum(t *testing.B) {
	for i := 0; i < t.N; i++ {
		stringSum()
	}
}
func BenchmarkTestStringSum2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		stringSum2()
	}
}
func BenchmarkTestStringSum3(t *testing.B) {
	for i := 0; i < t.N; i++ {
		stringSum3()
	}
}
