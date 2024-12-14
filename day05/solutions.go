package day05

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	return solvePart1(), solvePart2()
}

var rules [][]string
var pages [][]string

func solvePart1() int {
	sum := 0
	file, err := os.Open("day05/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "|") {
			nums := strings.Split(scanner.Text(), "|")
			rules = append(rules, nums)
		} else if strings.Contains(scanner.Text(), ",") {
			numbers := strings.Split(scanner.Text(), ",")
			pages = append(pages, numbers)
		}
	}

	for i := range len(pages) {
		flag := true
		for j := range len(pages[i]) - 1 {
			if !respectsRules(pages[i][j], pages[i][j+1]) {
				flag = false
				break
			}
		}
		if flag {
			middleStr := pages[i][len(pages[i])/2]
			middleInt, err := strconv.Atoi(middleStr)
			if err != nil {
				panic(err)
			}
			sum += middleInt
		}
	}

	return sum
}

func solvePart2() int {
	return 0
}

// Helpers
func respectsRules(a, b string) bool {
	for i := range len(rules) {
		//fmt.Printf("Checking %s and %s against %s\n", a, b, rules[i])
		hasA := rules[i][0] == a || rules[i][1] == a
		hasB := rules[i][0] == b || rules[i][1] == b
		if hasA && hasB {
			return rules[i][0] == a && rules[i][1] == b
		}
	}
	return true
}
