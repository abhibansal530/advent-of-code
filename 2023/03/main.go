package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"unicode"

	"github.com/abhibansal530/advent-of-code/utils"
)

func main() {
	file, err := os.Open("./input.txt")
	utils.PanicOnError(err)

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	answer := new(big.Int).SetUint64(0)
	starPosToNeighboursMap := make(map[int][]*big.Int)
	for rowNum := range grid {
		row := grid[rowNum]
		for col := 0; col < len(row); {
			digitIdx := strings.IndexFunc(row[col:], unicode.IsDigit)
			if digitIdx == -1 {
				// Done with current row.
				break
			}

			digitIdx += col
			nextIdx := strings.IndexFunc(row[digitIdx:], func(r rune) bool { return !unicode.IsDigit(r) })
			if nextIdx == -1 {
				nextIdx = len(row)
			} else {
				nextIdx += digitIdx
			}

			// Parse current number.
			var ok bool
			num := new(big.Int)
			num, ok = num.SetString(row[digitIdx:nextIdx], 10)
			if !ok {
				panic(row[digitIdx:nextIdx])
			}

			hasValidSymbolAsNeighbour := false
			cb := func(rowNum, colNum int) bool {
				x := grid[rowNum][colNum]
				if !unicode.IsDigit(rune(x)) && x != '.' {
					hasValidSymbolAsNeighbour = true
				}

				if x == '*' {
					pos := hashPosInGrid(rowNum, colNum, len(row))
					starPosToNeighboursMap[pos] = append(starPosToNeighboursMap[pos], num)
				}
				return true
			}

			visitNeighbours(grid, rowNum, digitIdx, nextIdx, cb)
			if hasValidSymbolAsNeighbour {
				// Is valid, add this number to our answer.
				answer = answer.Add(answer, num)
			}
			col = nextIdx
		}
	}

	// Find second part answer.
	answer2 := new(big.Int).SetUint64(0)
	for _, v := range starPosToNeighboursMap {
		if len(v) != 2 {
			continue
		}

		answer2 = answer2.Add(answer2, v[0].Mul(v[0], v[1]))
	}

	fmt.Printf("Answer for first part is %s\n", answer)
	fmt.Printf("Answer for second part is %s\n", answer2)
}

// Iterate on neighbours of number represented by grid[rowNum][startIndex..endIndex] using given
// callback. Iteration is stopped when callback returns false.
func visitNeighbours(grid []string, rowNum, startIndex, endIndex int, cb func(row, col int) bool) {
	neighbours := []struct {
		rowIdx, colFirstIdx, colLastIdx int
	}{
		// Previous row including diagonals.
		{
			rowIdx:      rowNum - 1,
			colFirstIdx: startIndex - 1,
			colLastIdx:  endIndex,
		},
		// Current row left.
		{
			rowIdx:      rowNum,
			colFirstIdx: startIndex - 1,
			colLastIdx:  startIndex - 1,
		},
		// Current row right.
		{
			rowIdx:      rowNum,
			colFirstIdx: endIndex,
			colLastIdx:  endIndex,
		},
		// Next row including diagonals.
		{
			rowIdx:      rowNum + 1,
			colFirstIdx: startIndex - 1,
			colLastIdx:  endIndex,
		},
	}

	for _, n := range neighbours {
		if n.rowIdx < 0 || n.rowIdx >= len(grid) {
			continue
		}

		row := grid[n.rowIdx]
		for j := n.colFirstIdx; j <= n.colLastIdx; j++ {
			if j < 0 || j >= len(row) {
				continue
			}

			if !cb(n.rowIdx, j) {
				return
			}
		}
	}
}

func hashPosInGrid(row, col, numCols int) int {
	return (row * numCols) + col
}
