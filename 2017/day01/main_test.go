package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	nums, err := parseInputString("1122")
	require.NoError(t, err)
	assert.Equal(t, 3, part1(nums))

	nums, err = parseInputString("1111")
	require.NoError(t, err)
	assert.Equal(t, 4, part1(nums))

	nums, err = parseInputString("1234")
	require.NoError(t, err)
	assert.Equal(t, 0, part1(nums))

	nums, err = parseInputString("91212129")
	require.NoError(t, err)
	assert.Equal(t, 9, part1(nums))
}

func TestPart2(t *testing.T) {
	nums, err := parseInputString("1212")
	require.NoError(t, err)
	assert.Equal(t, 6, part2(nums))

	nums, err = parseInputString("1221")
	require.NoError(t, err)
	assert.Equal(t, 0, part2(nums))

	nums, err = parseInputString("123425")
	require.NoError(t, err)
	assert.Equal(t, 4, part2(nums))

	nums, err = parseInputString("123123")
	require.NoError(t, err)
	assert.Equal(t, 12, part2(nums))

	nums, err = parseInputString("12131415")
	require.NoError(t, err)
	assert.Equal(t, 4, part2(nums))
}
