package common

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

// StringToInts takes in a string and seporator and returns a slce of ints
func StringToInts(in string, sep string) (out []int, err error) {
	for _, i := range strings.Split(in, sep) {
		num, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		out = append(out, num)
	}
	return out, nil
}

func DetectAnagram(word string, others []string) bool {
	runeMap := make(map[rune]int, len(word))
	for _, c := range word {
		runeMap[c]++
	}
	log.Println(runeMap)
	sort.Strings(others)
	for _, otherWord := range others {
		found := true
		if len(word) != len(otherWord) {
			// fast fail since it has to have all the letters
			continue
		}
		otherRunes := make(map[rune]int, len(otherWord))
		for _, c := range otherWord {
			otherRunes[c]++
		}

		for r, count := range otherRunes {
			origCount, ok := runeMap[r]
			// this is we found a letter in the other word not in the orig
			if !ok {
				found = false
			}
			// log.Printf("looking at [%v] comparing [%v] count [%v] to [%v]", otherRunes, r, count, origCount)

			if count != origCount {
				// we use a rune more than the sample word
				found = false
			}
		}
		return found
	}
	// We did not find any in the list of others
	return false
}
