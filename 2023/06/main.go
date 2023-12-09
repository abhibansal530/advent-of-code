package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/abhibansal530/advent-of-code/utils"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	utils.PanicOnError(err)

	lines := strings.Split(string(data), "\n")

	timePrefix := "Time:"
	timesStr := lines[0][len(timePrefix):]
	times := utils.ParseNumsFromStr(timesStr)

	distancePrefix := "Distance:"
	distancesStr := lines[1][len(distancePrefix):]
	distances := utils.ParseNumsFromStr(distancesStr)

	var answer int64 = 1
	for i := range times {
		answer = answer * numWaysToBeat(times[i], distances[i])
	}

	fmt.Printf("Answer for first part is: %d\n", answer)

	// Solve for second part.
	timeStr := strings.Join(strings.Fields(timesStr), "")
	distanceStr := strings.Join(strings.Fields(distancesStr), "")

	time, err := strconv.ParseInt(timeStr, 10, 64)
	utils.PanicOnError(err)
	distance, err := strconv.ParseInt(distanceStr, 10, 64)
	utils.PanicOnError(err)
	fmt.Printf("Answer for second part is: %d\n", numWaysToBeat(time, distance))
}

func numWaysToBeat(time, distance int64) int64 {
	// Possible ways :
	// (0 * (time - 0)), (1 * (time - 1)), (2 * (time - 2)),...., (time * (time - time)).
	//
	// In general :
	// i * (time - i), 0 <= i <= time
	//
	// So we want :
	// i * (time - i) > distance, or
	// i^2 - (i * time) + distance < 0
	//
	// Now if this quadratic equation has no roots, then answer is 0.
	//
	// Else if roots are r1 and r2, then simply count numbers in range (r1, r2).

	var a, b, c float64
	a = 1
	b = -1 * float64(time)
	c = float64(distance)

	d := math.Pow(b, 2) - (4 * a * c)
	if d < 0 {
		return 0
	}

	root := math.Sqrt(d)
	r1 := (-b + root) / (2 * a)
	r2 := (-b - root) / (2 * a)

	// Swap to have r1 <= r2.
	if r1 > r2 {
		r1, r2 = r2, r1
	}

	rc1 := math.Trunc(r1)
	if rc1 <= r1 {
		rc1 += 1
	}
	rc2 := math.Trunc(r2)
	if rc2 >= r2 {
		rc2 -= 1
	}

	// Answer will be non-negative values in [l, r].
	l := int64(rc1)
	r := int64(rc2)
	if r < 0 {
		return 0
	}

	l = max(0, l)
	if l > r {
		return 0
	}
	return r - l + 1
}
