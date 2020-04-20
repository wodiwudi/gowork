package leetcode

import (
	"fmt"
	"testing"
)

/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/
//思路：从根开始 如果为空节点返回nil 否则交换左右根地址，在递归左右子树。自顶向下的思想。
func TestTweentFour(t *testing.T) {
	arr := []int{-1, 2, -8, -10}
	fmt.Println(quickOrder(arr))
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

//选择排序
func chooseOrder(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
	return arr
}

//冒泡排序
func bubbleOrder(arr []int) []int {
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
func insertOrder(arr []int) []int {
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

//快速排序
func quickOrder(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	return quick(arr, 0, length-1)
}
func quick(arr []int, start, end int) []int {
	if start >= end {
		return nil
	}
	pivot := findPivot(arr, start, end)
	quick(arr, start, pivot-1)
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
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[pivot], arr[left] = arr[left], arr[pivot]
	return left
}
