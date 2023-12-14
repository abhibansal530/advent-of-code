package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/abhibansal530/advent-of-code/utils"
)

func main() {
	file, err := os.Open("./input.txt")
	utils.PanicOnError(err)

	var answer, answer2 int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := utils.ParseNumsFromStr(scanner.Text())
		x, y := findFirstAndNextVal(nums)
		answer += y
		answer2 += x
	}

	fmt.Printf("Answer for first part is: %d\n", answer)
	fmt.Printf("Answer for second part is: %d\n", answer2)
}

func findFirstAndNextVal(nums []int64) (int64, int64){
	var res1, res2 int64
	var mul int64 = 1
	for len(nums) > 0 && !hasOnlyZeros(nums) {
		res1 += nums[0] * mul
		res2 += nums[len(nums) - 1]
		nums = getDifferences(nums)
		mul *= -1
	}
	return res1, res2
}

func hasOnlyZeros(nums []int64) bool {
	for _, v := range nums {
		if v != 0 {
			return false
		}
	}
	return true
}

func getDifferences(nums []int64) []int64 {
	size := len(nums)
	for i := 0; i < size - 1; i++ {
		nums[i] = nums[i + 1] - nums[i]
	}
	return nums[:size - 1]
}
