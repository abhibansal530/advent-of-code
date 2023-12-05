package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abhibansal530/advent-of-code/utils"
)

var spelledDigits map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	file, err := os.Open("./input.txt")
	utils.PanicOnError(err)
	var totalFirstPart, totalSecondPart int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalFirstPart += int64(getCalibrationValue(line, false))
		totalSecondPart += int64(getCalibrationValue(line, true))
	}
	fmt.Printf("Answer for first part is %d\n", totalFirstPart)
	fmt.Printf("Answer for second part is %d\n", totalSecondPart)
}

func getCalibrationValue(line string, parseSpelledDigits bool) int {
	var first, last int
	first = -1

	cb := func(curr int) bool {
		if first == -1 {
			first = curr
		}
		last = curr
		return true
	}

	parseLine(line, parseSpelledDigits, cb)
	if first == -1 {
		return 0
	}

	return 10*first + last
}

// Parse given line in and invoke callback on found digits. Break parsing if callback returns false.
// parseSpelledDigits tells whether to parse spelled digits ("one", "two" etc.)
func parseLine(line string, parseSpelledDigits bool, callback func(currDigit int) bool) {
	for c := range line {
		x := line[c]
		if x >= '0' && x <= '9' {
			if res := callback(int(x - '0')); !res {
				return
			}
			continue
		}

		if !parseSpelledDigits {
			continue
		}

		// Try parsing spelled digits.
		dig, found := extractSpelledDigit(line, c)
		if found {
			if res := callback(dig); !res {
				return
			}
		}
	}
}

func extractSpelledDigit(input string, startIdx int) (int, bool) {
	for spelled, dig := range spelledDigits {
		if strings.HasPrefix(input[startIdx:], spelled) {
			return dig, true
		}
	}
	return 0, false
}
