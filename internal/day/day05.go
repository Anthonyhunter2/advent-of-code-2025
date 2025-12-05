package day

import (
	// "fmt"
	"strconv"
	"strings"
)

type Day5 struct{}

func init() {
	Days.RegisterDay(5, &Day5{})
}

func (d *Day5) SolvePart1(input []byte) (string, error) {
	freshIDs := make([][]int, 0)
	availableIDs := []int{}
	for line := range strings.SplitSeq(string(input), "\n") {
		if strings.Contains(line, "-") {
			freshRangeBegin, _ := strconv.Atoi(strings.Split(line, "-")[0])
			freshRangeEnd, _ := strconv.Atoi(strings.Split(line, "-")[1])
			x := []int{freshRangeBegin, freshRangeEnd}
			freshIDs = append(freshIDs, x)
		} else {
			if strings.TrimSpace(line) != "" {
				availableIntID, _ := strconv.Atoi(line)
				availableIDs = append(availableIDs, availableIntID)
			}
		}
	}
	availableFreshProducts := 0
	// rottenProducts := 0
	for i := 0; i < len(availableIDs); i++ {
		// freshProduct := availableFreshProducts
		for j := 0; j < len(freshIDs); j++ {
			if (freshIDs[j][0] <= availableIDs[i]) && (availableIDs[i] <= freshIDs[j][1]) {
				availableFreshProducts++
				break
			}
		}
		// if freshProduct == availableFreshProducts {
		// 	// fmt.Println(availableIDs[i], "is a rotten product")
		// 	rottenProducts++
		// }
	}
	return strconv.Itoa(availableFreshProducts), nil
}

func (d *Day5) SolvePart2(input []byte) (string, error) {
	freshIDs := make([][]int, 0)
	consolidatedFreshIDs := make([][]int, 0)
	// availableIDs := []int{}
	for line := range strings.SplitSeq(string(input), "\n") {
		if strings.Contains(line, "-") {
			freshRangeBegin, _ := strconv.Atoi(strings.Split(line, "-")[0])
			freshRangeEnd, _ := strconv.Atoi(strings.Split(line, "-")[1])
			x := []int{freshRangeBegin, freshRangeEnd}
			freshIDs = append(freshIDs, x)
		}
	}
	// Trying to create every number of this list is impossible, so now
	// we need to figure out which ranges overlap and where
	// we'll need to do the math on the uniq ranges to get our total number of fresh ids

	for i := 0; i < len(freshIDs); i++ {
		start := freshIDs[i][0]
		end := freshIDs[i][1]
		newStart := start
		newEnd := end
		for j := i + 1; j < len(freshIDs); j++ {
			checkStart := freshIDs[j][0]
			checkEnd := freshIDs[j][1]
			// if the list we're checking is already part of our list. make it 0's
			if (checkStart >= start) && (checkEnd <= end) {
				// setting that list as 0's since it doens't need to be included
				freshIDs[j][0] = 0
				freshIDs[j][1] = 0
				continue
			}
			// check to see if start or end is inbetween the ranges that follow
			if (checkStart <= start && start <= checkEnd) || (checkStart <= end && end <= checkEnd) {
				if start < checkStart {
					newStart = start
				} else {
					newStart = checkStart
				}
				if end > checkEnd {
					newEnd = end
				} else {
					newEnd = checkEnd
				}
				// we've moved the "i" range down to the "j" index.
				// we'll set it as 0's here
				freshIDs[i][0] = 0
				freshIDs[i][1] = 0
				// and reset j index to be the new total range
				freshIDs[j][0] = newStart
				freshIDs[j][1] = newEnd
			}
		}
		if start == newStart && end == newEnd {
			consolidatedFreshIDs = append(consolidatedFreshIDs, []int{newStart, newEnd})
		}

	}
	allFreshIDS := 0
	for i := 0; i < len(consolidatedFreshIDs); i++ {
		// we're now ignoreing anything we set to 0's
		if (consolidatedFreshIDs[i][1] == 0) && (consolidatedFreshIDs[i][0] == 0) {
			continue
		}
		// add the number of numbers in the range.. the +1 accounts for the fact that 3-5 has 3
		// numbers in it, not 2
		allFreshIDS += (consolidatedFreshIDs[i][1] - consolidatedFreshIDs[i][0]) + 1
	}

	return strconv.Itoa(allFreshIDS), nil
	// 436576159981259 - high
	// 352681648086138 - low
	// 352681648086146 - correct
}
