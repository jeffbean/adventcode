package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

const _dataFile = "day05/input.txt"

func readInJumps(data []byte) (a []int) {
	for _, line := range strings.Split(string(data), "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, num)
	}
	return a
}

func walkTheJumps(jumps []int) (numOfJumps int) {
	jumpNum := 0
	for i := 0; i < len(jumps); i += jumpNum {
		// log.Printf("begining jumpNum: %v - array %q ", jumpNum, jumps)
		jumpNum = jumps[i]
		jumps[i]++
		numOfJumps++
		if i+jumpNum < 0 {
			jumpNum = 0
		}
	}
	return numOfJumps
}

func walkTheJumps2(jumps []int) (resultingSlice []int, numOfJumps int) {
	jumpNum := 0
	for i := 0; i < len(jumps); i += jumpNum {
		// log.Printf("begining jumpNum: %v - array %v ", jumpNum, jumps)
		jumpNum = jumps[i]
		numOfJumps++
		if jumpNum >= 3 {
			jumps[i]--
		} else {
			jumps[i]++
		}
		if i+jumpNum < 0 {
			jumpNum = 0
		}
	}
	return jumps, numOfJumps
}

func part1() int {
	return walkTheJumps(readInJumps(common.ReadFile(_dataFile)))
}

func part2() int {
	_, n := walkTheJumps2(readInJumps(common.ReadFile(_dataFile)))
	return n

}

func main() {
	fmt.Printf("%v\n", part1())
	fmt.Printf("%v\n", part2())

}
