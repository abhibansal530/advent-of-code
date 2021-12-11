package main

import (
	"fmt"
	"github.com/abhibansal530/advent-of-code/utils"
)

func main() {
	input := utils.ReadFromFile("./input.txt")
	if input == nil {
		fmt.Println("Failed to read input.")
		return
	}


	fmt.Printf("Part A: %d\n", PartA(input))
	fmt.Printf("Part B: %d\n", PartB(input, 3))
}

func PartA(input []int) int {
	count := 0
	for idx := range input {
		if idx > 0 && input[idx] > input[idx - 1] {
			count = count + 1
		}
	}
	return count
}

func PartB(input []int, win_size int) int {
	count := 0
	prev, curr := 0, 0
	for idx := range input {
		if idx < win_size {
			curr = curr + input[idx]
			continue
		}

		prev = curr
		curr = curr + input[idx] - input[idx - win_size]
		if curr > prev {
			count = count + 1
		}
	}
	return count
}
