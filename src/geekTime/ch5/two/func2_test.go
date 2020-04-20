package two

import "testing"

func Sum(nums ...int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}

func TestMultiCanShu(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3))
}
