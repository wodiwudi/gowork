package leetcode

import (
	"testing"
)

/*
在n*n方阵里填入1,2,...,n*n,要求填成蛇形。例如n=4时方阵为：

10 11 12 1
9 16 13 2
8 15 14 3
7 6 5 4
*/
func TestThirtyFour(t *testing.T) {
	//num := 4
	getSnake(4)
}
func getSnake(num int) {
	//初始化
	result := [][]int{}
	for k := 0; k < num; k++ {
		tmp := make([]int, num)
		result = append(result, tmp)
	}
	//max := num * num
	count := 1
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			result[i][num-1-i] = count
		}
	}
}
