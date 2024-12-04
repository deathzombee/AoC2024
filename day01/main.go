package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/deathzombee/AoC2024/utils"
)

func readNumbers() ([]int, []int) {
	lines, _ := utils.ReadLines("input.txt")

	var leftSide, rightSide []int

	for _, line := range lines {
		nums := strings.Split(line, "   ")
		if len(nums) != 2 {
			fmt.Printf("Unexpected line format: %s\n", line)
			continue
		}

		left, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Printf("Error converting left number: %v\n", err)
			continue
		}
		leftSide = append(leftSide, left)

		right, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Printf("Error converting right number: %v\n", err)
			continue
		}
		rightSide = append(rightSide, right)
	}

	// Sort both sides
	sort.Ints(leftSide)
	sort.Ints(rightSide)

	return leftSide, rightSide
}

func pairDistance(left, right []int) int {
	total := 0

	// Ensure both slices have the same length
	if len(left) != len(right) {
		fmt.Println("Error: Left and right sides have different lengths")
		return 0
	}

	for i := 0; i < len(left); i++ {
		distance := abs(left[i] - right[i])
		total += distance
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func similarity(left, right []int) int {
	score := 0
	for _, l := range left {
		frequency := 0
		for _, r := range right {
			if l == r {
				frequency++
			}
		}
		score += l * frequency
	}
	return score
}

func part1() {
	left, right := readNumbers()
	distance := pairDistance(left, right)
	fmt.Println("Part 1 Result:", distance)
}

func part2() {
	left, right := readNumbers()
	score := similarity(left, right)
	fmt.Println("Part 2 Score:", score)
}

func main() {
	fmt.Println("Advent of Code 2024 - Day 1")

	fmt.Println("\nRunning Part 1:")
	part1()

	fmt.Println("\nRunning Part 2:")
	part2()
}
