package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
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

func part1(sliceOfNumbers []int) int {
	nextIndex := 0
	sumOfNumbers := 0

	for i, number := range sliceOfNumbers {
		nextIndex = i + 1
		if i == len(sliceOfNumbers)-1 {
			nextIndex = 0
		}
		log.Printf("number: %v, index: %v, nextIndex: %v, sum: %v", number, i, nextIndex, sumOfNumbers)
		if number == sliceOfNumbers[nextIndex] {
			sumOfNumbers += number
		}
	}
	return sumOfNumbers
}

func main() {

	nums, err := parseInputString(string(common.ReadFile("day01/input.txt")))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(nums))
}
