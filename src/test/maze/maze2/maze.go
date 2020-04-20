package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

//上左下右方向
var drivers = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) add(dot point) point {
	return point{p.i + dot.i, p.j + dot.j}
}

func (p point) at(arr [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(arr) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(arr[p.i]) {
		return 0, false
	}
	return arr[p.i][p.j], true
}

//走迷宫
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		current := Q[0]
		Q = Q[1:]
		if current == end {
			break
		}
		for _, driver := range drivers {
			next := current.add(driver)
			//判断越界和其他条件
			value, ok := next.at(maze)
			if !ok || value == 1 {
				continue
			}
			value, ok = next.at(steps)
			if !ok || value != 0 {
				continue
			}
			if next == start {
				continue
			}
			currentValue, _ := current.at(steps)
			steps[next.i][next.j] = currentValue + 1
			Q = append(Q, next)
		}
	}
	return steps
}

//找到最优解
func findLine() {

}

//走出迷宫的步数
func findMaxStep() {

}

func readfile(path string) [][]int {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	//初始化迷宫
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
	}
	for i := range maze {
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
func main() {
	//读文件
	maze := readfile("src/test/maze/maze.in")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%3d", steps[i][j])
		}
		fmt.Println()
	}
}
