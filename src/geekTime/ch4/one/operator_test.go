package two

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestOperator(t *testing.T) {
	a := 7 //0111
	i := a &^ Readable
	t.Log(i&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
