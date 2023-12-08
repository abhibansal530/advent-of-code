package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Read given file and parse each line as integer.
func ReadFromFile(filename string) (result []int) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Failed to read file: " + err.Error())
		return nil
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		result = append(result, number)
	}
	return result
}

// Parse "12 34 56" into []{12, 34, 56}.
func ParseNumsFromStr(s string) []int64 {
	var res []int64
	sl := strings.Fields(s)
	for _, x := range sl {
		y, err := strconv.ParseInt(x, 10, 64)
		PanicOnError(err)
		res = append(res, y)
	}
	return res
}
