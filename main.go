package main

import (
	"advent/day01"
	"advent/day02"
	"fmt"
)

func main() {
	day1Part1, day1Part2 := day01.Solve()
	fmt.Printf("Day 1:\n  Part 1: %v\n  Part 2: %v\n", day1Part1, day1Part2)

	day2Part1, day2Part2 := day02.Solve()
	fmt.Printf("Day 2:\n  Part 1: %v\n  Part 2: %v\n", day2Part1, day2Part2)
}
