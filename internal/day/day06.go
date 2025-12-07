package day

import (
	"fmt"
	"strconv"
	"strings"
)

type Day6 struct{}

func init() {
	Days.RegisterDay(6, &Day6{})
}

func (d *Day6) SolvePart1(input []byte) (string, error) {
	mathInput := make([][]string, 0)
	for line := range strings.SplitSeq(string(input), "\n") {
		fields := strings.Fields(line) // Splits on one or more whitespace characters
		mathInput = append(mathInput, fields)
	}

	grandTotal := 0

	for i := 0; i < len(mathInput[0]); i++ {
		total := 0
		for j := 0; j < len(mathInput)-1; j++ {
			if j == 0 {
				currentNumber, _ := strconv.Atoi(mathInput[j][i])
				total += currentNumber
				continue
			} else {
				currentNumber, _ := strconv.Atoi(mathInput[j][i])
				switch string(mathInput[len(mathInput)-1][i]) {
				case string('*'):
					total = total * currentNumber
				case string('+'):
					total = total + currentNumber
				}
			}
		}
		grandTotal += total
	}
	return strconv.Itoa(grandTotal), nil
}

func (d *Day6) SolvePart2(input []byte) (string, error) {
	mathInput := make([][]string, 0)
	for line := range strings.SplitSeq(string(input), "\n") {
		fields := strings.Split(line, "")
		mathInput = append(mathInput, fields)
	}
	fmt.Println(mathInput)
	return "", fmt.Errorf("not implemented")
}
