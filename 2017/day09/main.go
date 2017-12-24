package main

import (
	"fmt"

	"github.com/jeffbean/adventcode/2017/common"
)

func part1(data []byte) int {
	totalScore := 0
	groupScore := 0
	garbage := false
	for i := 0; i < len(data); i++ {
		r := data[i]
		switch r {
		case 33:
			i++
		case 60:
			garbage = true
		case 62:
			garbage = false
		case 123:
			if !garbage {
				groupScore++
			}
		case 125:
			if !garbage {
				totalScore += groupScore
				groupScore--
			}
		}
		// fmt.Printf("pos %v rune: %#U groupScore %v\n", i, r, groupScore)
	}
	return totalScore
}

func part2(data []byte) int {
	totalScore := 0
	garbage := false
	for i := 0; i < len(data); i++ {
		r := data[i]
		switch r {
		case 33:
			i++
			continue
		case 60:
			if garbage {
				totalScore++
			}
			garbage = true
		case 62:
			garbage = false
		default:
			if garbage {
				totalScore++
			}
		}
		// fmt.Printf("pos %v rune: %#U groupScore %v\n", i, r, groupScore)
	}
	return totalScore
}

func main() {
	data := common.GetInputFile()
	// log.Println(data)
	fmt.Printf("%v\n", part1(data))
	fmt.Printf("%v\n", part2(data))
}
