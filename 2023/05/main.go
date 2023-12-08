package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/abhibansal530/advent-of-code/utils"
)

type interval struct {
	start, end int64
}

type disjointIntervals []interval

type rangeEntry struct {
	source, dest, count int64
}

type rangeMap struct {
	ranges []rangeEntry
}

func main() {
	file, err := os.Open("./input.txt")
	utils.PanicOnError(err)

	var seeds []int64
	var allMaps []*rangeMap
	var currRangeMap *rangeMap

	scanner := bufio.NewScanner(file)
	for lineNum := 0; scanner.Scan(); lineNum++ {
		line := scanner.Text()
		if lineNum == 0 {
			// Parse seeds.
			prefix := "seeds: "
			seeds = utils.ParseNumsFromStr(line[len(prefix):])
			continue
		}

		if strings.Contains(line, "map") {
			// A new map is going to begin.
			//
			// 1. Store previous map if any.
			if currRangeMap != nil {
				allMaps = append(allMaps, currRangeMap)
			}

			// 2. Init new map.
			currRangeMap = new(rangeMap)
			continue
		}

		// Parse current line and add to currRangeMap.
		if currRangeMap != nil && len(line) > 0 {
			nums := utils.ParseNumsFromStr(line)
			utils.AssertEvalsToTrue(func() bool { return len(nums) == 3 })
			currRangeMap.ranges = append(currRangeMap.ranges, rangeEntry{
				dest:   nums[0],
				source: nums[1],
				count:  nums[2],
			})
		}
	}

	// Store last map.
	if currRangeMap != nil {
		allMaps = append(allMaps, currRangeMap)
	}

	// Find min seed location for first part.
	var answer int64 = math.MaxInt64
	for _, s := range seeds {
		answer = min(answer, findSeedLocation(s, allMaps))
	}
	fmt.Printf("Answer for first part is %d\n", answer)

	// Find min seed location for second part.
	var seedRanges []interval
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, interval{start: seeds[i], end: seeds[i] + seeds[i+1] - 1})
	}
	fmt.Printf("Answer for second part is %d\n", findMinSeedLocationsFromRanges(seedRanges, allMaps))
}

func findSeedLocation(seed int64, allMaps []*rangeMap) int64 {
	curr := seed
	for _, currMap := range allMaps {
		curr = lookupMap(curr, currMap)
	}
	return curr
}

func lookupMap(src int64, currMap *rangeMap) int64 {
	candidate := src
	for _, rng := range currMap.ranges {
		if src >= rng.source && src <= rng.source+rng.count {
			candidate = rng.dest + (src - rng.source)
			break
		}
	}
	return candidate
}
func findMinSeedLocationsFromRanges(seedIntervals []interval, allMaps []*rangeMap) int64 {
	curr := mergeIntervals(seedIntervals)
	for _, currMap := range allMaps {
		curr = lookupIntervals(curr, currMap)
	}

	var res int64 = math.MaxInt64
	for _, c := range curr {
		res = min(res, c.start)
	}
	return res
}

func lookupIntervals(input disjointIntervals, currMap *rangeMap) disjointIntervals {
	var res []interval
	for _, curr := range input {
		target := lookupInterval(curr, currMap)
		res = append(res, target...)
	}
	return mergeIntervals(res)
}

func lookupInterval(src interval, currMap *rangeMap) disjointIntervals {
	var res disjointIntervals
	currMap.ranges = sortRanges(currMap.ranges)
	for i := 0; i < len(currMap.ranges) && src.start <= src.end; {
		rng := currMap.ranges[i]
		rngStart, rngEnd := rng.source, rng.source+rng.count-1
		if rngEnd < src.start || rngStart > src.end {
			// No overlap, hence skip.
			i++
			continue
		}

		//   |-------------src------
		//  |---rng
		if rngStart <= src.start {
			res = append(res, interval{
				start: (src.start - rngStart) + rng.dest,
				end:   (min(src.end, rngEnd) - rngStart) + rng.dest,
			})
			src.start = rngEnd + 1
			i++
			continue
		}

		//   |-------------src------
		//     |---rng
		res = append(res, interval{start: src.start, end: rngStart - 1})
		src.start = rngStart
	}

	if src.start <= src.end {
		res = append(res, interval{start: src.start, end: src.end})
	}
	return mergeIntervals(res)
}

func mergeIntervals(input []interval) disjointIntervals {
	var res disjointIntervals
	if len(input) == 0 {
		return res
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i].start < input[j].start
	})

	curr := input[0]
	for i := 1; i < len(input); i++ {
		if input[i].start <= curr.end {
			// Merge this interval with running interval.
			curr.end = max(curr.end, input[i].end)
			continue
		}

		// Push curr and re-init curr.
		res = append(res, curr)
		curr = input[i]
	}

	res = append(res, curr)
	return res
}

func sortRanges(input []rangeEntry) []rangeEntry {
	sort.Slice(input, func(i, j int) bool {
		return input[i].source < input[j].source
	})
	return input
}
