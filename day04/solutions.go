package day04

import (
	"bufio"
	"os"
)

// 8 way adjacency for part 1
var directions = [][2]int{
	{0, 1},   // right
	{0, -1},  // left
	{1, 0},   // up
	{-1, 0},  // down
	{1, 1},   // up-right
	{1, -1},  // up-left
	{-1, 1},  // down-right
	{-1, -1}, // down-left
}

// Only diagonals for part 2
var diagonals = [][2]int{
	{-1, -1}, // top-left
	{-1, 1},  // top-right
	{1, -1},  // bottom-left
	{1, 1},   // bottom-right
}

func Solve() (int, int) {
	return solvePart1(), solvePart2()
}

func solvePart1() int {
	var grid [][]rune
	var count int
	file, err := os.Open("day04/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	// Convert input into [][]rune grid
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'X' {
				for _, dir := range directions {
					if checkDirection(grid, row, col, dir[0], dir[1]) {
						count++
					}
				}
			}
		}
	}

	return count
}

func solvePart2() int {
	var grid [][]rune
	var count int
	file, err := os.Open("day04/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	// Convert input into [][]rune grid
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'A' {
				// Only check in the forward directions from A
				topLeft := checkDiagonals(grid, row-1, col-1, 1, 1)
				topRight := checkDiagonals(grid, row-1, col+1, 1, -1)

				if topLeft && topRight {
					count++
				}
			}
		}
	}
	return count
}

// Helpers
func checkDirection(grid [][]rune, row, col int, colDir, rowDir int) bool {
	target := []rune{'M', 'A', 'S'}

	for i := 0; i < len(target); i++ {
		nextRow := row + (i+1)*rowDir
		nextCol := col + (i+1)*colDir

		if !inBounds(grid, nextRow, nextCol) || grid[nextRow][nextCol] != target[i] {
			return false
		}
	}

	return true
}

func checkDiagonals(grid [][]rune, row, col, dRow, dCol int) bool {
	if !inBounds(grid, row, col) ||
		!inBounds(grid, row+dRow, col+dCol) ||
		!inBounds(grid, row+2*dRow, col+2*dCol) {
		return false
	}

	start := grid[row][col]
	mid := grid[row+dRow][col+dCol]
	end := grid[row+2*dRow][col+2*dCol]

	// Either MAS or SAM sequence is valid
	return (start == 'M' && mid == 'A' && end == 'S') ||
		(start == 'S' && mid == 'A' && end == 'M')
}

func inBounds(grid [][]rune, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}
