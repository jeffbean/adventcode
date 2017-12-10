package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

var (
	partOneStarting = func([]int) int { return 1 }
	partTwoStarting = func(in []int) int { return len(in) / 2 }
)

func parseInputString(in string) (out []int, err error) {
	for _, i := range strings.Split(in, "") {
		num, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		out = append(out, num)
	}
	return out, nil
}

func findingSums(sliceOfNumbers []int, startingNextIndex func([]int) int) int {
	nextIndex := startingNextIndex(sliceOfNumbers)
	sumOfNumbers := 0

	for _, number := range sliceOfNumbers {
		if nextIndex == len(sliceOfNumbers) {
			nextIndex = 0
		}
		// log.Printf("%v, number: %v, index: %v, nextIndex: %v, sum: %v", sliceOfNumbers, number, i, nextIndex, sumOfNumbers)
		if number == sliceOfNumbers[nextIndex] {
			sumOfNumbers += number
		}
		nextIndex++
	}
	return sumOfNumbers
}

func part1(sliceOfNumbers []int) int {
	return findingSums(sliceOfNumbers, partOneStarting)
}

func part2(sliceOfNumbers []int) int {
	return findingSums(sliceOfNumbers, partTwoStarting)
}

func main() {
	nums, err := parseInputString(string(common.ReadFile("day01/input.txt")))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %v", part1(nums))
	fmt.Printf("part2: %v", part2(nums))
}
