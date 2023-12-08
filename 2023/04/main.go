package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/abhibansal530/advent-of-code/utils"
)

func main() {
	file, err := os.Open("./input.txt")
	utils.PanicOnError(err)

	scanner := bufio.NewScanner(file)
	var answer int64
	var scores []int
	for scanner.Scan() {
		line := scanner.Text()
		_, after, _ := strings.Cut(line, ":")
		winning, have, _ := strings.Cut(after, "|")
		winningNums := parseNums(winning)
		haveNums := parseNums(have)
		cnt := winningCount(haveNums, winningNums)
		scores = append(scores, cnt)
		if cnt > 0 {
			answer += int64(math.Pow(2, float64(cnt - 1)))
		}
	}

	fmt.Printf("Test %d\n", solve([]int{4, 2, 2, 1, 0, 0}))
	fmt.Printf("Answer for first part is %d\n", answer)
	fmt.Printf("Answer for second part is %d\n", solve(scores))
}

func parseNums(s string) []int {
	var res []int
	sl := strings.Fields(s)
	for _, x := range sl {
		y, err := strconv.Atoi(x)
		utils.PanicOnError(err)
		res = append(res, y)
	}
	return res
}

func winningCount(have, winning []int) int {
	var count int
	lookup := make(map[int]struct{})
	for _, x := range have {
		lookup[x] = struct{}{}
	}

	for _, x := range winning {
		if _, ok := lookup[x]; ok {
			count++
		}
	}
	return count
}

func solve(scores []int) int64 {
	updates := make([]int64, len(scores))
	var ans int64
	for i, s := range scores {
		// Update prefix sum.
		if i > 0 {
			updates[i] += updates[i - 1]
		}

		// We will get `updates[i] + 1` copy for each of `s` next cards.
		ans += int64(updates[i] + 1)
		if s == 0 {
			// No copies recieved.
			continue
		}

		if i + 1 < len(scores) {
			updates[i + 1] += updates[i] + 1
		}

		if i + s + 1 < len(scores) {
			updates[i + s + 1] -= (updates[i] + 1)
		}
	}
	return ans
}
