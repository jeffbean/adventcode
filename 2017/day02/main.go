package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

const _spreadSheetSep = "\t"

var dataFile = "day02/input.txt"

func readInSpreadsheet(data []byte) (ss [][]int) {
	for _, line := range strings.Split(string(data), "\n") {
		nums, err := common.StringToInts(line, _spreadSheetSep)
		if err != nil {
			log.Fatal(err)
		}
		ss = append(ss, nums)
	}
	return ss
}

func checkSumMatrix(ss [][]int) int {
	var checkSum int
	for row, b := range ss {
		var min = b[0]
		var max = b[0]
		for _, c := range b {
			if c < min {
				min = c
			}
			if c > max {
				max = c
			}
		}
		diff := max - min
		checkSum += diff
		log.Printf("row %v, min: %v, max: %v, diff: %v", row, min, max, diff)
	}
	return checkSum
}

func checksumDivides(ss [][]int) int {
	var checkSum int
	for row, b := range ss {
		var divisor int
		for i, c := range b {
			for j := 0; j < len(b); j++ {
				if i == j {
					continue
				}
				r := c % b[j]
				if r == 0 {
					divisor = c / b[j]
					break
				}
			}
		}

		checkSum += divisor
		log.Printf("row %v, divisor: %v", row, divisor)
	}
	return checkSum
}

func part1(ss [][]int) int {
	return checkSumMatrix(ss)
}

func part2(ss [][]int) int {
	return checksumDivides(ss)
}

func main() {
	data := common.ReadFile(dataFile)

	fmt.Printf("part1 answer: %v\n", part1(readInSpreadsheet(data)))
	fmt.Printf("part2 answer: %v\n", part2(readInSpreadsheet(data)))

}
