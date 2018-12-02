package day02

import (
	"testing"

	"github.com/jeffbean/adventcode/2018/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		assert.Equal(t, 12, part1([]string{
			"abcdef",
			"bababc",
			"abbcde",
			"abcccd",
			"aabcdd",
			"abcdee",
			"ababab",
		}))
	})

	t.Run("input test", func(t *testing.T) {
		input, err := input.FileToLines("input.txt")
		require.NoError(t, err, "failed to read in input")
		num := part1(input)
		t.Log(num)
	})
}
