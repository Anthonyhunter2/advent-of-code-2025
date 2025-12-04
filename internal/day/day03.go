package day

import (
	"strconv"
	"strings"
)

type Day3 struct{}

func init() {
	Days.RegisterDay(3, &Day3{})
}

func (d *Day3) SolvePart1(input []byte) (string, error) {
	jibawatts := 0
	for line := range strings.SplitSeq(string(input), "\n") {
		// battery, _ := strconv.Atoi(line)
		bat_max := 0
		str_bat_max := ""
		max_index := 0
		second_largest := 0
		str_second_largest := ""
		// The len(line) -1 is because our max can't be the last number of the battery
		for index := 0; index < len(line)-1; index++ {
			bat_value, _ := strconv.Atoi(string(line[index]))
			// fmt.Println(index, bat_value)
			if bat_value > bat_max {
				max_index = index
				bat_max = bat_value
				str_bat_max = string(line[index])
			}
		}
		for i := (max_index + 1); i < len(line); i++ {
			sec_value, _ := strconv.Atoi(string(line[i]))
			// fmt.Println("sec_value ", second_largest, "i value", i)
			if sec_value > second_largest {
				second_largest = sec_value
				str_second_largest = string(line[i])
			}
		}
		new_jibbybits, _ := strconv.Atoi(str_bat_max + str_second_largest)
		// fmt.Println(new_jibbybits)
		jibawatts += new_jibbybits
	}
	return strconv.Itoa(jibawatts), nil
}

func (d *Day3) SolvePart2(input []byte) (string, error) {
	jibawatts := 0
	// newJibbybits := ""
	for line := range strings.SplitSeq(string(input), "\n") {
		newJibbybits := ""
		indexPoint := 0
		for i := 11; i >= 0; i-- {
			nextDigit, lastIndex := findNextDigit(line, i, indexPoint)
			newJibbybits = newJibbybits + strconv.Itoa(nextDigit)
			indexPoint = lastIndex + 1
		}
		jibbybitsInt, _ := strconv.Atoi(newJibbybits)
		jibawatts += jibbybitsInt
	}
	return strconv.Itoa(jibawatts), nil
}

func findNextDigit(numList string, numberLeft int, startingIndex int) (int, int) {
	// were finding the next available digit with room to find
	// the remaining numbers we need.
	highestAvailable := 0
	highestIndex := 0
	numList = numList[startingIndex:]
	for i := 0; i < len(numList)-numberLeft; i++ {
		numValue, _ := strconv.Atoi(string(numList[i]))
		if numValue > highestAvailable {
			highestAvailable = numValue
			highestIndex = i
		}
	}
	return highestAvailable, highestIndex + startingIndex
}
