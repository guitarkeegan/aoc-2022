package main

import (
	"fmt"
	"os"
	"strings"
)

var dbg = func() func(format string, as ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(format string, as ...any) {}
	}
	return func(format string, as ...any) {
		fmt.Printf(format+"\n", as...)
	}
}()

type Matrix [][]string

func (m Matrix) String() string {
	display := ""
	for i := range m {
		for j := range m[0] {
			display += m[i][j]
		}
		display += "\n"
	}
	return display
}

func makeMatrix(input string) Matrix {
	matrix := [][]string{}
	lines := strings.SplitSeq(input, "\n") // Fixed: use Split instead of SplitSeq
	for line := range lines {
		if line != "" { // Skip empty lines
			matrix = append(matrix, strings.Split(line, ""))
		}
	}
	return matrix
}

func getStartPos(m Matrix) (int, int) {
	for i := range m {
		for j := range m[0] {
			if m[i][j] == "S" {
				dbg("Start pos: row: %d, col: %d", i, j)
				return i, j
			}
		}
	}
	dbg("S not found")
	return -1, -1
}

func runTheRun(m Matrix) int {
	shortest := 9999999999
	for i := range m {
		for j := range m[0] {
			if m[i][j] == "S" || m[i][j] == "a" {
				shortest = min(shortest, Run2(m, i, j))
			}
		}
	}
	return shortest
}

func Run(m Matrix, r, c int) int {

	var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// state
	minSteps := 99999
	visited := make(map[[2]int]struct{})
	visited[[2]int{r, c}] = struct{}{}

	var bfs func()
	bfs = func() {
		ROWS := len(m)
		COLS := len(m[0])
		q := [][3]int{{r, c, 0}}

		for len(q) > 0 {

			cur := q[0]
			dbg("r: %d, c: %d", cur[0], cur[1])
			q = q[1:]
			for _, dir := range dirs {

				nxtRow, nxtCol := cur[0]+dir[0], cur[1]+dir[1]
				// out of bounds
				if nxtRow < 0 ||
					nxtCol < 0 ||
					nxtRow >= ROWS ||
					nxtCol >= COLS {
					continue
				}
				// visited
				if _, ok := visited[[2]int{nxtRow, nxtCol}]; ok {
					continue
				}
				if m[cur[0]][cur[1]] == "S" {
					if "a"[0]+1 < m[nxtRow][nxtCol][0] {
						continue
					}
					visited[[2]int{nxtRow, nxtCol}] = struct{}{}
					q = append(q, [3]int{nxtRow, nxtCol, cur[2] + 1})
					continue
				}
				// reached the end
				if m[nxtRow][nxtCol] == "E" {
					dbg("found E")
					if m[cur[0]][cur[1]][0]+1 < "z"[0] {
						continue
					}
					minSteps = min(minSteps, cur[2]+1)
					q = [][3]int{}
					break
				}
				// invalid move
				if m[cur[0]][cur[1]][0]+1 < m[nxtRow][nxtCol][0] {
					dbg("invalid move")
					continue
				}
				visited[[2]int{nxtRow, nxtCol}] = struct{}{}
				q = append(q, [3]int{nxtRow, nxtCol, cur[2] + 1})
			}
		}
	}

	bfs()
	return minSteps

}

func Run2(m Matrix, r, c int) int {

	var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// state
	minSteps := 99999
	visited := make(map[[2]int]struct{})
	visited[[2]int{r, c}] = struct{}{}

	var bfs func()
	bfs = func() {
		ROWS := len(m)
		COLS := len(m[0])
		q := [][3]int{{r, c, 0}}

		for len(q) > 0 {

			cur := q[0]
			dbg("r: %d, c: %d", cur[0], cur[1])
			q = q[1:]
			for _, dir := range dirs {

				nxtRow, nxtCol := cur[0]+dir[0], cur[1]+dir[1]
				// out of bounds
				if nxtRow < 0 ||
					nxtCol < 0 ||
					nxtRow >= ROWS ||
					nxtCol >= COLS {
					continue
				}
				// visited
				if _, ok := visited[[2]int{nxtRow, nxtCol}]; ok {
					continue
				}
				// reached the end
				if m[nxtRow][nxtCol] == "E" {
					dbg("found E")
					if m[cur[0]][cur[1]][0]+1 < "z"[0] {
						continue
					}
					minSteps = min(minSteps, cur[2]+1)
					q = [][3]int{}
					break
				}
				// invalid move
				if m[cur[0]][cur[1]][0]+1 < m[nxtRow][nxtCol][0] {
					dbg("invalid move")
					continue
				}
				visited[[2]int{nxtRow, nxtCol}] = struct{}{}
				q = append(q, [3]int{nxtRow, nxtCol, cur[2] + 1})
			}
		}
	}

	bfs()
	return minSteps

}

func main() {

	m := makeMatrix(input)
	// r, c := getStartPos(m)
	// fmt.Printf("Matrix:\n%s", m)
	//
	// fmt.Println(Run(m, r, c))
	fmt.Println(runTheRun(m))
}
