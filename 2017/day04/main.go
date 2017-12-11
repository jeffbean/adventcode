package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

const _part1Sep = " "

var dataFile = "day04/input.txt"

type validator func(string) bool

var validationPart1 = []validator{
	validateDuplicaes,
}

var validationPart2 = []validator{
	validateDuplicaes,
	validateAnagram,
}

func parseFileLine(data []byte) []string {
	return strings.Split(string(data), "\n")
}

func validateDuplicaes(phrase string) bool {
	themWords := make(map[string]struct{})
	words := strings.Split(phrase, _part1Sep)
	for _, w := range words {
		if _, ok := themWords[w]; ok {
			return false
		}
		themWords[w] = struct{}{}
	}
	return true
}

func validateAnagram(phrase string) bool {
	origWords := strings.Split(phrase, _part1Sep)
	sort.Strings(origWords)
	// fmt.Printf("origwords: %v\n", origWords)
	for i := 0; i < len(origWords); i++ {
		others := make([]string, len(origWords))
		copy(others, origWords)
		copy(others[i:], others[i+1:])
		others[len(others)-1] = ""
		others = others[:len(others)-1]
		if common.DetectAnagram(origWords[i], others) {
			log.Printf("anagram: %q -> %v", origWords[i], others)
			return false
		}
	}
	return true
}

func validPassphrases(lines []string, vs []validator) int {
	validLines := 0
	for _, line := range lines {
		validLine := true
		for _, v := range vs {
			if !v(line) {
				validLine = false
				continue
			}
		}
		if validLine {
			log.Printf("valid line:   %q", line)
			validLines++
		} else {
			log.Printf("invalid line: %q", line)
		}
	}
	return validLines
}

func part1() int {
	return validPassphrases(parseFileLine(common.ReadFile(dataFile)), validationPart1)
}

func part2() int {
	return validPassphrases(parseFileLine(common.ReadFile(dataFile)), validationPart2)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
