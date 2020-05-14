package leetcode

import (
	"fmt"
	"testing"
)

/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
示例 1:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
*/

//动态规划，每一个格子都存储为以当前i,j结尾的最大数值
//如果i，j都为0说明是第一个元素 跳过 如果i为0则为第一列元素，只能向下走，
//如果j为0则为第一行元素,只能向右走。否则为比较向下和向右谁大就走哪里
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += getMax(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[m-1][n-1]
}

func TestFourTyEight(t *testing.T) {
	arrs := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(maxValue(arrs))
}
