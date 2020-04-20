package error

import (
	"errors"
	"testing"
)

func Fibo(n int) ([]int, error) {
	if n < 2 || n > 10000 {
		return nil, errors.New("n 需要在2到10000之间")
	}
	result := []int{1, 1}
	for i := 2; i < n; i++ {
		result = append(result, result[i-1]+result[i-2])
	}
	return result, nil
}

func TestFibo(t *testing.T) {
	if nums, err := Fibo(10); err != nil {
		t.Error(err)
	} else {
		t.Log(nums)
	}
}
