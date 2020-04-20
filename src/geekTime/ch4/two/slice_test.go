package two

import "testing"

func TestSlice(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSlice2(t *testing.T) {
	s := []string{"q", "b", "c", "h", "o", "p"}
	sb := s[1:3]
	t.Log(len(sb), cap(sb))
	sb[0] = "xxx"
	t.Log(s)
}
