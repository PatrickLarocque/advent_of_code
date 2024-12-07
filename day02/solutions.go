package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	return solvePart1(), solvePart2()
}

func solvePart1() int {
	// Number of 'safe' reports
	var count int

	file, err := os.Open("day02/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	// Loop over ever line in input.txt
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		nums := make([]int, len(strs))

		// convert each []string into []int
		for i := range len(strs) {
			num, err := strconv.Atoi(strs[i])
			if err != nil {
				panic(err)
			}
			nums[i] = num
		}

		// True = descending, false = ascending
		descending := nums[0] > nums[1]
		// Flip to false if we ever fail the satisfiedSafeCondition for given subarray
		safe := true

		// verify each []int subarray against 'safe' condition
		for i := range len(nums) - 1 {
			diff := nums[i] - nums[i+1]
			if !satisfiesSafeConditions(diff, descending) {
				safe = false
			}
		}
		if safe {
			count++
		}
	}
	return count
}

func solvePart2() int {
	return 0
}

// Helpers
func satisfiesSafeConditions(diff int, descending bool) bool {
	if descending && diff <= 0 {
		return false
	}

	if !descending && diff >= 0 {
		return false
	}

	if getAbsoluteDifference(diff) > 3 {
		return false
	}

	return true
}

func getAbsoluteDifference(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
