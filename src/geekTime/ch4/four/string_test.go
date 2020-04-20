package four

import (
	"strconv"
	"strings"
	"testing"
)

func TestUnicodeAndUtf8(t *testing.T) {
	s := "中"
	ru := []rune(s)
	by := []byte(s)
	t.Logf("中 unicode is %x", ru[0]) //rune 为unicode  string和byte为utf8
	t.Logf("中 utf8 is %x", s)
	t.Logf("中 utf8 is %x", by[0])

	s2 := "中国人"
	for _, v := range s2 { //range 和 字符串迭代输出的是rune不使byte
		t.Logf("%[1]d,%[1]c", v) //[1]代表都用1
	}

	s3 := "abcdefg"
	for _, v := range s3 { //range 和 字符串迭代输出的是rune不使byte
		t.Logf("%[1]d,%[1]c", v) //[1]代表都用1
	}
}

func TestStrings(t *testing.T) {
	s := "A:B:C"
	parts := strings.Split(s, ":")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))

	var a = 10
	s = strconv.Itoa(a)
	t.Logf("%[1]T,%[1]s", s)
	if v, err := strconv.Atoi(s); err == nil {
		t.Log(a + v)
	}
}
