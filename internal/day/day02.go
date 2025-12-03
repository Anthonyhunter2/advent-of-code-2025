package day

import (
	"reflect"
	"strconv"
	"strings"
)

type Day2 struct{}

func init() {
	Days.RegisterDay(2, &Day2{})
}

func (d *Day2) SolvePart1(input []byte) (string, error) {
	total_of_item_ids := 0
	item_id_range := strings.Split(string(input), ",")
	for _, id_range := range item_id_range {
		range_start, _ := strconv.Atoi(strings.Split(id_range, "-")[0])
		range_end, _ := strconv.Atoi(strings.Split(id_range, "-")[1])
		for item_id := range_start; item_id <= range_end; item_id++ {
			if HasRepeatingSequence(item_id) {
				// fmt.Printf("%v\n", item_id)
				total_of_item_ids += item_id
			}

		}
	}
	return strconv.Itoa(total_of_item_ids), nil
}

func (d *Day2) SolvePart2(input []byte) (string, error) {
	total_of_item_ids := 0
	item_id_range := strings.Split(string(input), ",")
	for _, id_range := range item_id_range {
		range_start, _ := strconv.Atoi(strings.Split(id_range, "-")[0])
		range_end, _ := strconv.Atoi(strings.Split(id_range, "-")[1])
		for item_id := range_start; item_id <= range_end; item_id++ {
			num_seq := []int{}
			str_item_id := strconv.Itoa(item_id)
			for i := range str_item_id {
				num_seq = append(num_seq, int(str_item_id[i]))
			}
			if HasRepeatingSequence2(num_seq) {
				// fmt.Printf("%v\n", item_id)
				total_of_item_ids += item_id
			}

		}
	}
	return strconv.Itoa(total_of_item_ids), nil
}

func HasRepeatingSequence(item_id int) bool {
	idString := strconv.Itoa(item_id)
	idLength := len(idString)

	// if the id is a single digit, it is not invalid
	if idLength == 1 {
		return false
	}

	part1 := idString[:idLength/2]
	part2 := idString[idLength/2:]

	// check for leading zeros
	if part1[0] == 0 || part2[0] == 0 {
		return false
	}

	return part1 == part2
}

func HasRepeatingSequence2(sequence []int) bool {
	n := len(sequence)
	if n < 2 {
		return false // A sequence needs at least two elements to potentially repeat
	}

	// Iterate through possible pattern lengths
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		isRepeating := true
		// Compare sub-sequences
		for i := patternLen; i < n; i += patternLen {
			// Get the current pattern to compare
			currentPattern := sequence[i-patternLen : i]

			// Get the next segment to compare against
			end := i + patternLen
			if end > n {
				end = n // Handle cases where the last segment is shorter
			}
			nextSegment := sequence[i:end]

			// If lengths don't match or content differs, it's not a repeating pattern of this length
			if len(currentPattern) != len(nextSegment) || !reflect.DeepEqual(currentPattern, nextSegment) {
				isRepeating = false
				break
			}
		}
		if isRepeating {
			return true // Found a repeating pattern
		}
	}
	return false // No repeating pattern found
}
