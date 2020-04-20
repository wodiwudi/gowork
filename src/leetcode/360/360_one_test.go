package main

import (
	"fmt"
	"testing"
)

func winner(human, win int, arr []int) int {
	if len(arr) != human {
		return -1
	}
	m := make(map[int]int, 0)
	for _, v := range arr {
		m[v] = 0
	}
	result := 0
	for {
		if arr[0] > arr[1] {
			m[arr[0]]++
			tmp := arr[1]
			arr = append(arr[0:1], arr[2:]...)
			arr = append(arr, tmp)
			if m[arr[0]] >= win {
				result++
				break
			}
		} else {
			m[arr[1]]++
			arr = append(arr[1:], arr[0])
			if m[arr[1]] >= win {
				result++
				break
			}
		}
		result++
	}
	return result
}
func TestWinner(t *testing.T) {
	human := 4
	win := 3
	arr := []int{1, 3, 2, 4}
	//arr2 := []int{3, 2, 4, 1}
	fmt.Println(winner(human, win, arr))
}
