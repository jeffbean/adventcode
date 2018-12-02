package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	lines, err := FileToLines("inputpart1")
	if err != nil {
		return err
	}
	// fmt.Printf("%v", input)
	ints, err := parseInput(lines)
	if err != nil {
		return err
	}
	fmt.Println(part2(ints))
	return nil
}

func FileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func parseInput(b []string) ([]int, error) {
	inInts := []int{}
	for _, s := range b {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		inInts = append(inInts, i)
	}
	return inInts, nil
}

func part1(input []int) int {
	var sum int
	for _, i := range input {
		sum = sum + i
	}
	return sum
}

func part2(input []int) int {
	tracking := []int{0}
	var frequency int
	var i int
	var answer int
	dup := false

	for !dup {
		// loop back through
		if i >= len(input) {
			i = 0
		}

		frequency = frequency + input[i]

		for _, j := range tracking {
			if frequency == j {
				dup = true
				answer = frequency
				break
			}
		}
		tracking = append(tracking, frequency)

		i++
	}
	return answer
}
