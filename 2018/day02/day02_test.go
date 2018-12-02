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
func TestPart2(t *testing.T) {
	t.Run("test single diffs", func(t *testing.T) {
		assert.Equal(t, true, diffStrings("fghij", "fguij"))
		assert.Equal(t, false, diffStrings("abcde", "axcye"))
		assert.Equal(t, true, diffStrings("abcdef", "bbcdef"))
		assert.Equal(t, false, diffStrings("abcdef", "bbcdeff"))
		assert.Equal(t, false, diffStrings("abcdefvvv", "abcdef"))
		assert.Equal(t, false, diffStrings("abcdef", "abcdefvvv"))
	})

	t.Run("example", func(t *testing.T) {
		assert.Equal(t, "fgij", part2([]string{
			"abcde",
			"fghij",
			"klmno",
			"pqrst",
			"fguij",
			"axcye",
			"wvxyz",
		}))
	})

	t.Run("input test", func(t *testing.T) {
		input, err := input.FileToLines("input.txt")
		require.NoError(t, err, "failed to read in input")
		result := part2(input)
		t.Log(result)
	})
}
