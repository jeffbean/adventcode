package day02

import (
	"math"
)

func part1(in []string) int {
	var (
		totalTwice  int
		totalThrice int
	)
	for _, s := range in {
		tw, th := processID(s)
		totalTwice += tw
		totalThrice += th
	}

	return totalTwice * totalThrice
}

func processID(id string) (twice, thrice int) {
	counters := make(map[rune]int)

	for _, c := range id {
		counters[c]++
	}

	for _, num := range counters {
		switch num {
		case 2:
			twice = 1
		case 3:
			thrice = 1
		}
	}

	return twice, thrice
}

// The boxes will have IDs which differ by exactly one character at the same position in both strings
func part2(ids []string) string {
	for _, id := range ids {
		for i := 0; i < len(ids); i++ {
			if diffStrings(id, ids[i]) {
				return stripUncommon(id, ids[i])
			}
		}
	}
	return ""
}

func diffStrings(left, right string) bool {
	diffs := math.Abs(float64(len(left) - len(right)))

	for i, r := range left {
		if i > len(right) || diffs > 1 {
			return false
		}
		if rune(right[i]) != r {
			diffs++
		}
	}

	return diffs == 1
}

func stripUncommon(left, right string) string {
	newString := []rune{}
	for i, r := range left {
		if rune(right[i]) == r {
			newString = append(newString, r)
		}
	}

	return string(newString)
}
