package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

const _part1Sep = " "

var dataFile = "day04/input.txt"

func parseFileLine(data []byte) []string {
	return strings.Split(string(data), "\n")
}

func validatePassphrase(phrase string) bool {
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
func validPassphrases(lines []string) int {
	validLines := 0
	for _, line := range lines {
		if validatePassphrase(line) {
			log.Printf("valid line: %q", line)
			validLines++
		}
	}
	return validLines
}

func part1() int {
	return validPassphrases(parseFileLine(common.ReadFile(dataFile)))
}

func main() {
	fmt.Println(part1())
}
