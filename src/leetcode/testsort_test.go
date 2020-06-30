package leetcode

import (
	"fmt"
	"testing"
)

//选择排序
func choosesort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
	return arr
}

//归并排序
func combinesort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	return combinemerge(combinesort(arr[:middle]), combinesort(arr[middle:]))
}

func combinemerge(leftArr, rightArr []int) []int {
	i, j := 0, 0
	lenleft, lenright := len(leftArr), len(rightArr)
	res := []int{}
	for i < lenleft && j < lenright {
		if leftArr[i] > rightArr[j] {
			res = append(res, rightArr[j])
			j++
			continue
		}
		res = append(res, leftArr[i])
		i++
	}
	res = append(res, leftArr[i:]...)
	res = append(res, rightArr[j:]...)
	return res
}

func TestTestSort(t *testing.T) {
	arr := []int{1, 4, 2, 7, 5, -1, 0}
	fmt.Println(combinesort(arr))
}
