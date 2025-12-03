package day

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day1 struct{}

func init() {
	Days.RegisterDay(1, &Day1{})
}

func (d *Day1) SolvePart1(input []byte) (string, error) {
	Dial := 50
	zeros := 0
	for line := range strings.SplitSeq(string(input), "\n") {
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("invalid distance given: %v", err)
		}
		if distance > 100 {
			distance = distance % 100
		}
		switch direction {
		case 'L':
			Dial -= distance
			if Dial < 0 {
				Dial += 100
			}
		case 'R':
			Dial += distance
			if Dial >= 100 {
				Dial -= 100
			}
		}
		if Dial == 0 {
			zeros++
		}
	}
	return fmt.Sprintf("%d", zeros), nil
}

func (d *Day1) SolvePart2(input []byte) (string, error) {
	Dial := 50
	zeros := 0
	for line := range strings.SplitSeq(string(input), "\n") {
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("invalid distance given: %v", err)
		}
		if distance > 100 {
			zeros += int(math.Floor(float64(distance) / 100))
			distance = distance % 100
		}

		switch direction {
		case 'L':
			if Dial == 0 {
				Dial += (100 - distance)
			} else {
				Dial -= distance
				if Dial <= 0 {
					zeros++
					if Dial != 0 {
						Dial += 100
					}
				}
			}
		case 'R':
			Dial += distance
			if Dial > 99 {
				Dial -= 100
				zeros++
			}
		}
		// fmt.Printf("zeros: %v, Dial: %v Distance Moved: %v\n", zeros, Dial, distance)

	}
	return fmt.Sprintf("%d", zeros), nil
}
