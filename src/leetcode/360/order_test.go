package main

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	arr := []int{-10, 2, -8, 1, 12, 5}
	fmt.Println(quickOrder(arr))
}
func chooseOrder(arr []int) []int {
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
func bubbleOrder(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func insertOrder(arr []int) []int {
	for i := 0; i < len(arr); i++ {
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

func combineOrder(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	return merge(combineOrder(arr[:middle]), combineOrder(arr[middle:]))
}
func merge(leftArr, rightArr []int) []int {
	i, j := 0, 0
	lenleft, lenright := len(leftArr), len(rightArr)
	result := []int{}
	for i < lenleft && j < lenright {
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

func quickOrder(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	return quick(arr, 0, length-1)
}
func quick(arr []int, start, end int) []int {
	if start > end {
		return nil
	}
	pivot := findPivot(arr, start, end)
	quick(arr, 0, pivot-1)
	quick(arr, pivot+1, end)
	return arr
}
func findPivot(arr []int, start, end int) int {
	pivot := start
	left, right := start, end
	for left != right {
		for left < right && arr[right] > arr[pivot] {
			right--
		}
		for left < right && arr[left] <= arr[pivot] {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[left], arr[right]
		}
	}
	arr[left], arr[start] = arr[start], arr[left]
	return left
}
