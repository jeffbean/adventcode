package common

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// ReadFile just reads a file and fatals if anything goes wrong.
func ReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	check(err)
	return data
}

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
