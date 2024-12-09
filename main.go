package main

import (
	"advent/day01"
	"advent/day02"
	"advent/day03"
	"advent/day04"
	"advent/day05"
	"fmt"
)

func main() {
	day1Part1, day1Part2 := day01.Solve()
	fmt.Printf("Day 1:\n  Part 1: %v\n  Part 2: %v\n", day1Part1, day1Part2)

	day2Part1, day2Part2 := day02.Solve()
	fmt.Printf("Day 2:\n  Part 1: %v\n  Part 2: %v\n", day2Part1, day2Part2)

	day3Part1, day3Part2 := day03.Solve()
	fmt.Printf("Day 3:\n  Part 1: %v\n  Part 2: %v\n", day3Part1, day3Part2)

	day4Part1, day4Part2 := day04.Solve()
	fmt.Printf("Day 4:\n  Part 1: %v\n  Part 2: %v\n", day4Part1, day4Part2)

	day5Part1, day5Part2 := day05.Solve()
	fmt.Printf("Day 5:\n  Part 1: %v\n  Part 2: %v\n", day5Part1, day5Part2)
}
