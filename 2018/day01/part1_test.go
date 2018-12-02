package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	assert.Equal(t, part1([]int{1, 1, 1}), 3)
	assert.Equal(t, part1([]int{1, 1, -2}), 0)
	assert.Equal(t, part1([]int{-1, -2, -3}), -6)
}

func TestRun(t *testing.T) {
	assert.NoError(t, run())
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 0, part2([]int{1, -1}))
	assert.Equal(t, 10, part2([]int{3, 3, 4, -2, -4}))

	assert.Equal(t, 5, part2([]int{-6, 3, 8, 5, -6}))
}
