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
