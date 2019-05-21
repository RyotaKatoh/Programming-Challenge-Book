package main

import "fmt"

var field [][]string = [][]string{
	{"W", ".", ".", ".", ".", ".", ".", ".", ".", "W", "W", "."},
	{".", "W", "W", "W", ".", ".", ".", ".", ".", "W", "W", "W"},
	{".", ".", ".", ".", "W", "W", ".", ".", ".", "W", "W", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", "W", "W", "."},
	{".", ".", ".", ".", ".", ".", ".", ".", ".", "W", ".", "."},
	{".", ".", "W", ".", ".", ".", ".", ".", ".", "W", ".", "."},
	{".", "W", ".", "W", ".", ".", ".", ".", ".", "W", "W", "."},
	{"W", ".", "W", ".", "W", ".", ".", ".", ".", ".", "W", "."},
	{".", "W", ".", "W", ".", ".", ".", ".", ".", ".", "W", "."},
	{".", ".", "W", ".", ".", ".", ".", ".", ".", ".", "W", "."},
}

var ly int = len(field)
var lx int = len(field[0])

func dfs(x, y int) {
	field[y][x] = "."

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if x+dx >= 0 && x+dx < lx && y+dy >= 0 && y+dy < ly && field[y+dy][x+dx] == "W" {
				dfs(x+dx, y+dy)
			}
		}
	}

}

func solve() int {
	ans := 0

	for y := 0; y < ly; y++ {
		for x := 0; x < lx; x++ {
			if field[y][x] == "W" {
				ans++
				dfs(x, y)
			}
		}
	}

	return ans
}

func main() {
	// fmt.Println(field)
	fmt.Println(solve())

}
