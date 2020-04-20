package main

import (
	"fmt"
	"os"
)

//抽象出点位
type point struct {
	i, j int
}

//定义方向 上左下右
var directs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

//判断是否越界（最外围一圈）
func (p point) at(grip [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grip) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grip[p.i]) {
		return 0, false
	}
	return grip[p.i][p.j], true
}

//读取迷宫文件
func readMaze(path string) [][]int {
	file, err := os.Open(path) //打开文件
	if err != nil {
		panic(err)
	}
	//定义行和列
	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d", &row, &col) //6和5传给row和col
	maze := make([][]int, row)                   //maze是一个类型为[]int的切片，length为行数(整体代入的思想)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, _ = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

//point的加方法
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//start end起终点
func walk(maze [][]int, start, end point) [][]int {
	//初始化steps
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//建立队列，储存要探索的点位
	Q := []point{start}
	//队列不空就一直走下去，去发现
	for len(Q) > 0 {
		//初始为第一个
		current := Q[0]
		//剔除头
		Q = Q[1:]
		//如果遇到终点则退出
		if current == end {
			break
		}
		for _, dir := range directs {
			//下一点等于当前节点加方向
			next := current.add(dir)
			/*条件
			1.迷宫的下一个节点必须为0，1为墙
			2.steps的next也是0，如果有值则为走过了
			3.避免走回起点，next != start
			4.不能越界
			*/
			maze_at, ok := next.at(maze)
			//如果点位越界，或者迷宫的点位为强
			if !ok || maze_at == 1 {
				continue //跳过并继续循环
			}

			steps_at, ok := next.at(steps)
			//如果点位越界，或者走过此点位
			if !ok || steps_at != 0 {
				continue
			}
			//不能走回起点
			if next == start {
				continue
			}
			//开始探索迷宫 1.步数加一 2.将要探索的点放入队列
			curSteps, _ := current.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}
	return steps
}

//找到出迷宫的步数
func findStepInMaze(steps [][]int) int {
	max := 0
	for i := range steps {
		for j := range steps[i] {
			if steps[i][j] > max {
				max = steps[i][j]
			}
		}
	}
	return max
}

//找到迷宫的路线
func findLine(steps [][]int, end, start point) [][]int {
	//初始化二维数组作为路线图
	lines := make([][]int, len(steps))
	for i := range lines {
		lines[i] = make([]int, len(steps[i]))
	}
	Q := []point{end}

	for len(Q) > 0 {
		current := Q[0]
		Q = Q[1:]
		if current == start {
			break
		}
		//开始探索
		for _, direct := range directs {
			next := current.add(direct)
			//判断条件
			next_point, ok := next.at(steps)
			current_point, _ := current.at(steps)
			if !ok || next_point != current_point-1 {
				continue
			}
			//开始探索
			lines[current.i][current.j] = current_point
			lines[next.i][next.j] = next_point
			Q = append(Q, next)
		}
	}

	return lines
}

//广度优先算法走迷宫
func main() {
	maze := readMaze("src/test/maze/maze.in")

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	//找出步数
	stepInMaze := findStepInMaze(steps)
	fmt.Println("走出迷宫的步数为")
	fmt.Println(stepInMaze)

	//找出路线
	lines := findLine(steps, point{len(maze) - 1, len(maze[0]) - 1}, point{0, 0})
	fmt.Println("迷宫路线为：")
	for _, line := range lines {
		for _, point := range line {
			fmt.Printf("%2d ", point)
		}
		fmt.Println()
	}
}
