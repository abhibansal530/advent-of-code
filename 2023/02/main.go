package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abhibansal530/advent-of-code/utils"
)

type set struct {
	red, blue, green int
}

type game struct {
	id   int
	sets []set
}

func main() {
	file, err := os.Open("./input.txt")
	utils.PanicOnError(err)
	scanner := bufio.NewScanner(file)
	initial := set{red: 12, blue: 14, green: 13}
	var answer, powerSum int64
	for scanner.Scan() {
		line := scanner.Text()
		g := parseLine(line)
		if isPossible(initial, g) {
			answer += int64(g.id)
		}
		powerSum += minPower(g)
	}
	fmt.Printf("Answer for first part is: %d\n", answer)
	fmt.Printf("Answer for second part is: %d\n", powerSum)
}

// Example input -
// Game 1: 2 red, 2 green; 6 red, 3 green; 2 red, 1 green, 2 blue; 1 red
func parseLine(line string) game {
	var res game
	before, after, _ := strings.Cut(line, ": ")
	res.id, _ = cutInNumAndStr(before, false)

	allSets := strings.Split(after, "; ")
	for _, st := range allSets {
		res.sets = append(res.sets, parseSet(st))
	}
	return res
}

// Example input -
// 2 red, 2 green
func parseSet(s string) set {
	var res set
	got := strings.Split(s, ", ")
	for _, ball := range got {
		cnt, b := cutInNumAndStr(ball, true)
		switch b {
		case "red":
			res.red = cnt
		case "blue":
			res.blue = cnt
		case "green":
			res.green = cnt
		default:
			panic(fmt.Sprintf("invalid color: %s", b))
		}
	}
	return res
}

func cutInNumAndStr(s string, numComesFirst bool) (int, string) {
	before, after, _ := strings.Cut(s, " ")
	if numComesFirst {
		before, after = after, before
	}

	num, err := strconv.Atoi(after)
	utils.PanicOnError(err)
	return num, before
}

func isPossible(total set, g game) bool {
	for _, s := range g.sets {
		if s.red > total.red || s.blue > total.blue || s.green > total.green {
			return false
		}
	}
	return true
}

func minPower(g game) int64 {
	var red, blue, green int64
	for _, st := range g.sets {
		red = max(red, st.red)
		blue = max(blue, st.blue)
		green = max(green, st.green)
	}
	return red * blue * green
}

func max(x int64, y int) int64 {
	if x > int64(y) {
		return x
	}
	return int64(y)
}
