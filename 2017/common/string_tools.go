package common

import (
	"sort"
	"strconv"
	"strings"
)

// StringToInts takes in a string and seporator and returns a slce of ints
func StringToInts(in string, sep string) (out []int, err error) {
	for _, i := range strings.Split(in, sep) {
		num, err := strconv.Atoi(strings.TrimSpace(i))
		if err != nil {
			return nil, err
		}
		out = append(out, num)
	}
	return out, nil
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func DetectAnagram(words []string) bool {
	sort.Strings(words)
	otherWords := make(map[string]struct{})
	for _, otherWord := range words {
		otherWord = SortString(otherWord)
		// log.Printf("sorted word: %q -> %v", otherWord, otherWords)
		if _, ok := otherWords[otherWord]; ok {
			return true
		}
		otherWords[otherWord] = struct{}{}
	}
	return false
}
