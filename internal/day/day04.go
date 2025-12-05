package day

import (
	"strconv"
	"strings"
)

type Day4 struct{}

func init() {
	Days.RegisterDay(4, &Day4{})
}

func (d *Day4) SolvePart1(input []byte) (string, error) {
	coords := make([][]string, 0)
	for line := range strings.SplitSeq(string(input), "\n") {
		x := make([]string, len(line))
		parts := strings.Split(line, "")
		copy(x, parts)
		coords = append(coords, x)
	}
	accessablePaper := 0
	for x := 0; x < len(coords); x++ {
		for y := 0; y < len(coords[x]); y++ {
			// fmt.Println("current coords: ", x, y, coords[x][y])
			if coords[x][y] == "." {
				continue
			}
			if paperAccessable(coords, x, y) {
				// fmt.Println("Found available paper at:", x, y)
				accessablePaper++
			}

		}
	}
	return strconv.Itoa(accessablePaper), nil
}

func (d *Day4) SolvePart2(input []byte) (string, error) {
	coords := make([][]string, 0)
	for line := range strings.SplitSeq(string(input), "\n") {
		x := make([]string, len(line))
		parts := strings.Split(line, "")
		copy(x, parts)
		coords = append(coords, x)
	}
	totPaperRemoved := 0

	for {
		previousPaperRemoved := totPaperRemoved
		for x := 0; x < len(coords); x++ {
			for y := 0; y < len(coords[x]); y++ {
				// fmt.Println("current coords: ", x, y, coords[x][y])
				if coords[x][y] == "." {
					continue
				}
				if paperAccessable(coords, x, y) {
					// fmt.Println("Found available paper at:", x, y)
					totPaperRemoved++
					coords[x][y] = "."
				}

			}
		}
		if totPaperRemoved == previousPaperRemoved {
			break
		}
	}
	return strconv.Itoa(totPaperRemoved), nil
}

func meatAndPotatoes(coords [][]string) ([][]string, int) {
	paperRemoved := 0
	for x := 0; x < len(coords); x++ {
		for y := 0; y < len(coords[x]); y++ {
			// fmt.Println("current coords: ", x, y, coords[x][y])
			if coords[x][y] == "." {
				continue
			}
			if paperAccessable(coords, x, y) {
				// fmt.Println("Found available paper at:", x, y)
				paperRemoved++
				coords[x][y] = "."
			}

		}
	}
	return coords, paperRemoved
}

func paperAccessable(coords [][]string, x int, y int) bool {
	// all available coordinates around x and y.... ie (x-1,y-1)(x0,y-1)
	// blanks := 0
	paper := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			xCoordToCheck := x + i
			yCoordToCheck := y + j
			//checking out of bounds gives a blank space

			if xCoordToCheck >= 0 && xCoordToCheck < len(coords) && yCoordToCheck >= 0 && yCoordToCheck < len(coords[x]) && coords[xCoordToCheck][yCoordToCheck] == "@" {
				paper++
			}
			// if xCoordToCheck < 0 || yCoordToCheck < 0 || xCoordToCheck > len(coords)-1 || yCoordToCheck > len(coords[x])-1 {
			// 	blanks++
			// 	continue
			// }
			// fmt.Println("checking coords for position", x, y, "thigns", xCoordToCheck, yCoordToCheck, coords[xCoordToCheck][yCoordToCheck])
			// if coords[xCoordToCheck][yCoordToCheck] == "@" {
			// 	paper++
			// } else {
			// 	blanks++
			// }
			if paper == 4 {
				return false
			}
			// if blanks == 5 {
			// 	return true
			// }
		}
	}
	return true
}
