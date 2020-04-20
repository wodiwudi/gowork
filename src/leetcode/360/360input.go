package main

import (
	"fmt"
	"math"
)

/*
有一个城市需要修建，给你N个民居的坐标X,Y，问把这么多民居全都包进城市的话，城市所需最小面积是多少
（注意，城市为平行于坐标轴的正方形）
输入描述:
第一行为N，表示民居数目（2≤N≤1000）

输出描述:
城市所需最小面积
示例1
输入
2
0 0
2 2
输出
4
示例2
输入
2
0 0
0 3
输出
9
*/
func main() {
	human := 0
	win := 0
	_, _ = fmt.Scanf("%d%d", &human, &win)
	points := []int{}
	for i := 0; i < human; i++ {
		x := 0
		_, _ = fmt.Scanf("%d", &x)
		points = append(points, x)
	}
	fmt.Println(points)
}
func getMax(points [][]int) int {
	maxX, maxY, minX, minY := 0, 0, points[0][0], points[0][1]
	for _, v := range points {
		if v[0] < minX {
			minX = v[0]
		}
		if v[1] < minY {
			minY = v[0]
		}
		if v[0] > maxX {
			maxX = v[0]
		}
		if v[1] > maxY {
			maxY = v[1]
		}
	}
	max := int(math.Max(float64(maxX-minX), float64(maxY-minY)))
	return max * max
}
