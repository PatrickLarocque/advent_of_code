package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	return solvePart1(), solvePart2()
}

func solvePart1() int {
	value := readInput()
	fmt.Printf("%v\n", value)
	return 0
}

func solvePart2() int {
	return 0
}

// Helpers
func readInput() map[int][]int {
	value := make(map[int][]int)
	file, err := os.Open("day07/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		leftNum, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		rightNums := strings.Fields(parts[1])
		for _, numStr := range rightNums {
			num, _ := strconv.Atoi(strings.TrimSpace(numStr))
			value[leftNum] = append(value[leftNum], num)
		}
	}
	return value
}
