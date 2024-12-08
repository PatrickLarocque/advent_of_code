package day03

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func Solve() (int, int) {
	return solvePart1(), solvePart2()
}

func solvePart1() int {
	var sum int

	file, err := os.Open("day03/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	mulExpression := regexp.MustCompile(`mul\(([1-9]\d{0,2}),([1-9]\d{0,2})\)`)

	for scanner.Scan() {
		matches := mulExpression.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])

			if err1 != nil || err2 != nil {
				panic("An error occurred")
			}

			sum += num1 * num2
		}
	}

	return sum
}

func solvePart2() int {
	return 0
}
