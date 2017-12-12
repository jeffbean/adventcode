package common

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

const _inputFileName = "input.txt"

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

// GetInputFile reads in the eecutables input.txt file
func GetInputFile() []byte {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("caller filename not found")
	}
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}

	return ReadFile(filepath.Join(dir, _inputFileName))
}
