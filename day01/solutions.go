package day01

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

var leftArr, rightArr []int

func Solve() (int, int) {
	return solvePart1(), solvePart2()
}

func solvePart1() int {
	var sum int

	file, err := os.Open("day01/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		left, _ := strconv.Atoi(nums[0])
		right, _ := strconv.Atoi(nums[1])
		leftArr = append(leftArr, left)
		rightArr = append(rightArr, right)
	}
	sort.Ints(leftArr)
	sort.Ints(rightArr)

	for i := range len(leftArr) {
		diff := leftArr[i] - rightArr[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	return sum
}

func solvePart2() int {
	sum := 0
	for i := range len(leftArr) {
		freq := 0
		for j := range len(rightArr) {
			if leftArr[i] == rightArr[j] {
				freq++
			}
		}
		sum += leftArr[i] * freq
	}
	return sum
}
