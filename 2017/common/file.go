package common

import (
	"io/ioutil"
	"log"
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
