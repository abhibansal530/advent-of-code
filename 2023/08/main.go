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

	var instructions string
	nodesMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for lineNum := 0; scanner.Scan(); lineNum++ {
		if lineNum == 0 {
			instructions = scanner.Text()
			continue
		}

		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		// Example line -
		// LRL = (MCG, TRC)
		curr := line[:3]
		left := line[len("ABC = ("):][:3]
		right := line[len("ABC = (XYZ, "):][:3]
		nodesMap[curr] = []string{left, right}
	}

	var steps int64
	curr, target := "AAA", "ZZZ"
	for i := 0; curr != target; i++ {
		steps++
		dir := instructions[i%len(instructions)]
		child, ok := nodesMap[curr]
		if !ok {
			break
		}
		if dir == 'L' {
			curr = child[0]
		} else {
			curr = child[1]
		}
	}
	fmt.Printf("Answer for first part is %d\n", steps)

	var currNodes []string
	for k := range nodesMap {
		if k[2] == 'A' {
			currNodes = append(currNodes, k)
		}
	}

	/*
		// For curr = currNodes[i], it seems that there is only one node
		// ending with a 'Z' reachable from given curr. And also steps are
		// multiple of first distance. So our answer will be LCM of steps for
		// all sources.
		var steps int64
		curr := currNodes[1]
		for i := 0; ; i++ {
			if curr[2] == 'Z' {
				fmt.Printf("%s %d\n", curr, steps2)
			}
			steps2++
			dir := instructions[i%len(instructions)]
			if dir == 'L' {
				curr = nodesMap[curr][0]
			} else {
				curr = nodesMap[curr][1]
			}
		}
	*/

	var answer2 int64 = 1
	for i := range currNodes {
		curr := currNodes[i]
		var steps int64
		for j := 0; curr[2] != 'Z'; j++ {
			steps++
			dir := instructions[j%len(instructions)]
			if dir == 'L' {
				curr = nodesMap[curr][0]
			} else {
				curr = nodesMap[curr][1]
			}
		}
		answer2 = lcm(answer2, steps)
	}

	fmt.Printf("Answer for second part is %d\n", answer2)
}

func destReached(currNodes []string) bool {
	for _, n := range currNodes {
		if n[2] != 'Z' {
			return false
		}
	}
	return true
}

func lcm(a, b int64) int64 {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
