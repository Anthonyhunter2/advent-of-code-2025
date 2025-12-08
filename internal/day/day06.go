package day

import (
	"regexp"
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
	grandTotal := 0
	mathInput := make([][]string, 0)
	for line := range strings.SplitSeq(string(input), "\n") {
		x := make([]string, len(line))
		parts := strings.Split(line, "")
		copy(x, parts)
		mathInput = append(mathInput, x)
	}
	re := regexp.MustCompile(`\S`)
	splitLines := re.FindAllStringIndex(strings.Join(mathInput[len(mathInput)-1], ""), -1)
	for x := 0; x < len(splitLines); x++ {
		newInput := make([][]string, 6)
		startingOffset := 0
		endingOffset := 0
		for y := 0; y < len(mathInput); y++ {
			curNumber := []string{}
			if y == len(mathInput)-1 {
				operatorIndex := splitLines[x][0]
				newInput[5] = strings.Fields(mathInput[y][operatorIndex])
				continue
			}
			if x == len(splitLines)-1 {
				startingOffset = splitLines[x][0]
				curNumber = mathInput[y][startingOffset:]
			} else {
				startingOffset = splitLines[x][0]
				endingOffset = splitLines[x+1][0] - 1
				curNumber = mathInput[y][startingOffset:endingOffset]
			}
			for z := 0; z < len(curNumber); z++ {
				newInput[z] = append(newInput[z], curNumber[z])
			}
		}
		total := 0
		for i := 0; i < len(newInput); i++ {
			if len(newInput[i]) != 0 {
				numberString := strings.Fields(strings.Join(newInput[i], ""))
				numberInt, _ := strconv.Atoi(numberString[0])
				if i == 0 {
					total += numberInt
				} else {
					switch string(newInput[5][0]) {
					case string('*'):
						total = total * numberInt
					case string('+'):
						total = total + numberInt
					}
				}
			} else {
				grandTotal += total
				break
			}
		}

	}
	return strconv.Itoa(grandTotal), nil
}
