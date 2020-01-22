package main

import (
	"fmt"
	"os"
)

//--------------------
// 广度优先搜索算法走迷宫
//--------------------

// 获取迷宫矩阵
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	// Fscanf() 扫描文件获取文件内容
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point {
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

// 坐标值相加
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 当前所有坐标点
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

// 破解迷宫之 广度优先算法
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// 继续探索的条件
			// maze at next is 0，
			// and steps at next is 0，
			// and next != start
			// 才能继续探索
			val, ok := next.at(maze)
			if !ok || val == 1 {	// 越界
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {	// 已经走过的地方
				continue
			}
			if next == start {		// 走回了原点
				continue
			}

			// 将步数写进格子中
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			// 将要继续探索的位置放进队列里
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
