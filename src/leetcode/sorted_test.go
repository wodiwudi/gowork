package leetcode

import (
	"fmt"
	"testing"
)

//排序算法
func TestSorted(t *testing.T) {
	nums := []int{5, 1, -1, -8, -10, 9, 0, 11}
	fmt.Println(quickOrders(nums))
}

//选择排序
func chooseOrders(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

//冒泡排序
func bubbleOrders(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//插入排序
func selectOrders(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		pre := i - 1
		current := arr[i]
		for pre >= 0 && arr[pre] > current {
			arr[pre+1] = arr[pre]
			pre--
		}
		arr[pre+1] = current
	}
	return arr
}

//归并排序
func combineOrders(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	return merges(combineOrders(arr[:middle]), combineOrders(arr[middle:]))
}
func merges(leftArr, rightArr []int) []int {
	i, j := 0, 0
	lenLeft, lenRight := len(leftArr), len(rightArr)
	result := []int{}
	for i < lenLeft && j < lenRight {
		if leftArr[i] > rightArr[j] {
			result = append(result, rightArr[j])
			j++
			continue
		}
		result = append(result, leftArr[i])
		i++
	}
	result = append(result, leftArr[i:]...)
	result = append(result, rightArr[j:]...)
	return result
}

//快速排序
func quickOrders(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	return quicks(arr, 0, length-1)
}
func quicks(arr []int, start, end int) []int {
	if start > end {
		return nil
	}
	pivot := findPivots(arr, start, end)
	quicks(arr, start, pivot-1)
	quicks(arr, pivot+1, end)
	return arr
}
func findPivots(arr []int, start, end int) int {
	pivot := arr[start]
	left, right := start, end
	for left != right {
		for left < right && arr[right] > pivot {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	fmt.Println("pivot :", pivot, "arr[start] :", arr[start], "left", left)
	arr[start], arr[left] = arr[left], arr[start]
	return left
}
