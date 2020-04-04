package main


import (
	"fmt"
	"os"
)

//读取输入文件
func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("read file error!")
	}

	var row, col int
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

type Point struct {
	x, y int   //行、列
}

func (p Point) add(dir Point) Point {
	return Point{p.x + dir.x, p.y + dir.y}
}

func (p Point) at(grid [][]int) (int, bool) {
	if p.x < 0 || p.x >= len(grid) {
		return 0, false
	}

	if p.y < 0 || p.y >= len(grid[p.x]) {
		return 0, false
	}


	return grid[p.x][p.y], true
}

// 定义方向： 上左下右       {-1,0}代表上，表示行号减一，列号不变
var dirs = [4]Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start Point, end Point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	queue := []Point{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			break
		}

		for _, dir := range dirs {
			next := curr.add(dir)

			//maze at next is 0
			//and steps at next is 0
			//and next != start
			val, ok := next.at(maze)
			if !ok || val == 1{
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			currSteps, _ := curr.at(steps)
			steps[next.x][next.y] = currSteps + 1

			queue = append(queue, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")

	fmt.Println(len(maze))
	for _, row := range maze {
		for _, val := range row {
			fmt.Print(val, "	")
		}
		fmt.Println()
	}

	steps := walk(maze, Point{0, 0}, Point{len(maze)-1, len(maze[0])-1})
	fmt.Println("Steps:")
	for _, row := range steps {
		for _, val := range row {
			fmt.Print(val, "	")
		}
		fmt.Println()
	}
}

/*
package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
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

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

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

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] =
				curSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")

	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	// TODO: construct path from steps
}*/