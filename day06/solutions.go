package day06

import (
	"bufio"
	"os"
)

var directions = map[rune][]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

var rightTurn = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

var grid [][]rune

func Solve() (int, int) {
	readInput()
	return solvePart1(), solvePart2()
}

func solvePart1() int {
	var startRow, startCol int
	found := false
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '^' {
				startRow, startCol = row, col
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	// Create visited grid
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	currentRow, currentCol := startRow, startCol
	direction := '^'
	visited[currentRow][currentCol] = true

	for {
		nextRow := currentRow + directions[direction][0]
		nextCol := currentCol + directions[direction][1]

		// Check bounds
		if isOutOfBounds(grid, nextRow, nextCol) {
			break
		}

		if grid[nextRow][nextCol] == '#' { // Turn right
			visited[currentRow][currentCol] = true
			direction = rightTurn[direction]
		} else { // Keep going in the same direction
			currentRow = nextRow
			currentCol = nextCol
			visited[currentRow][currentCol] = true
		}
	}

	// Count visited positions
	count := 0
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] {
				count++
			}
		}
	}

	return count
}

type State struct {
	row, col int
	dir      rune
}

func solvePart2() int {
	// Find starting position with '^'
	var startRow, startCol int
	found := false
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '^' {
				startRow, startCol = row, col
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	cycles := 0

	// Try placing # at each position
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			// Skip the starting position
			if row == startRow && col == startCol {
				continue
			}

			// Save original value at this position
			originalValue := grid[row][col]

			// Place a # at current position
			grid[row][col] = '#'

			// Reset state for new attempt
			visited := make(map[State]bool)
			currentRow, currentCol := startRow, startCol
			direction := '^'

			// Run the navigation logic
			for {
				nextRow := currentRow + directions[direction][0]
				nextCol := currentCol + directions[direction][1]

				if isOutOfBounds(grid, nextRow, nextCol) {
					break
				}

				if isCycle(visited, nextRow, nextCol, direction) {
					cycles++
					break
				}

				if grid[nextRow][nextCol] == '#' { // Turn right
					visited[State{row: currentRow, col: currentCol, dir: direction}] = true
					direction = rightTurn[direction]
				} else { // Keep going in the same direction
					visited[State{row: nextRow, col: nextCol, dir: direction}] = true
					currentRow = nextRow
					currentCol = nextCol
				}
			}

			// Restore the original value at this position
			grid[row][col] = originalValue
		}
	}

	return cycles
}

// Helpers
func isOutOfBounds(grid [][]rune, nextRow, nextCol int) bool {
	return nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0])
}

func isCycle(visited map[State]bool, nextRow, nextCol int, dir rune) bool {
	nextState := State{
		row: nextRow,
		col: nextCol,
		dir: dir,
	}
	return visited[nextState]
}

// Read and make the grid
func readInput() {
	file, err := os.Open("day06/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

}
