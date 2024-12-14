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
	sum := 0
	for i := range len(pages) {
		flag := true
		for j := range len(pages[i]) - 1 {
			if !respectsRules(pages[i][j], pages[i][j+1]) {
				flag = false
				break
			}
		}
		if !flag {
			for k := range len(pages[i]) - 1 {
				for l := k + 1; l < len(pages[i]); l++ {
					pages[i] = checkAndSwap(pages[i][k], pages[i][l], pages[i], k, l)
				}
			}
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

// Helpers
func respectsRules(a, b string) bool {
	for i := range len(rules) {
		hasA := rules[i][0] == a || rules[i][1] == a
		hasB := rules[i][0] == b || rules[i][1] == b
		if hasA && hasB {
			return rules[i][0] == a && rules[i][1] == b
		}
	}
	return true
}

func checkAndSwap(a, b string, arr []string, idx1, idx2 int) []string {
	for i := range len(rules) {
		hasA := rules[i][0] == a || rules[i][1] == a
		hasB := rules[i][0] == b || rules[i][1] == b
		if hasA && hasB && !(rules[i][0] == a && rules[i][1] == b) {
			temp := arr[idx1]
			arr[idx1] = arr[idx2]
			arr[idx2] = temp
		}
	}
	return arr
}
