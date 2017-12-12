package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJumps(t *testing.T) {
	fileContents := []byte("1\n2\n3")
	ints := readInJumps(fileContents)
	assert.Equal(t, ints, []int{1, 2, 3})
}

func TestPart1(t *testing.T) {
	testJumps := []int{0, 3, 0, 1, -3}
	numOfJumps := walkTheJumps(testJumps)
	assert.Equal(t, numOfJumps, 5)

}

func TestPart2(t *testing.T) {
	testJumps := []int{0, 3, 0, 1, -3}
	a, numOfJumps := walkTheJumps2(testJumps)
	assert.Equal(t, numOfJumps, 10)
	assert.Equal(t, a, []int{2, 3, 2, 3, -1})

}
