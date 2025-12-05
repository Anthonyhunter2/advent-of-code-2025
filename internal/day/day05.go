package day

import (
	"fmt"
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
	// TODO: Implement part 2
	return "", fmt.Errorf("not implemented")
}
