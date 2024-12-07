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

		if isArraySafe(nums) {
			count++
		}
	}
	return count
}

func solvePart2() int {
	count := 0
	file, err := os.Open("day02/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		nums := make([]int, len(strs))

		// Convert strings to ints
		for i := range strs {
			num, err := strconv.Atoi(strs[i])
			if err != nil {
				panic(err)
			}
			nums[i] = num
		}

		if isArraySafeWithProblemDampener(nums) {
			count++
		}
	}
	return count
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

func isArraySafeWithProblemDampener(nums []int) bool {
	// First check if the array is safe without removing any elements
	if isArraySafe(nums) {
		return true
	}

	// Try removing each element one at a time
	for i := range nums {
		// Create a new array without the current element
		dampened := make([]int, 0, len(nums)-1)
		dampened = append(dampened, nums[:i]...)
		dampened = append(dampened, nums[i+1:]...)

		// Check if the dampened array is safe
		if isArraySafe(dampened) {
			return true
		}
	}

	return false
}

func isArraySafe(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	// Determine initial direction based on first two elements
	descending := nums[0] > nums[1]

	// Check each pair of adjacent numbers
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i] - nums[i+1]

		// Check both direction consistency and difference magnitude
		if !satisfiesSafeConditions(diff, descending) {
			return false
		}
	}

	return true
}
