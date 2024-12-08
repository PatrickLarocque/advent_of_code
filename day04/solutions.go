package day04

import (
	"bufio"
	"os"
)

var directions = [][2]int{
	{0, 1},   // right
	{0, -1},  // left
	{1, 0},   // up
	{-1, 0},  //down
	{1, 1},   // up-right
	{1, -1},  // up-left
	{-1, 1},  // down-right
	{-1, -1}, // down-left
}

var grid [][]rune

func Solve() (int, int) {
	return solvePart1(), solvePart2()
}

func solvePart1() int {
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
	return 0
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

func inBounds(grid [][]rune, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}
