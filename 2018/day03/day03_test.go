package day03

import (
	"testing"

	"github.com/jeffbean/adventcode/2018/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		assert.Equal(t, 4, part1([]string{
			"#1 @ 1,3: 4x4",
			"#2 @ 3,1: 4x4",
			"#3 @ 5,5: 2x2",
		}))
	})

	t.Run("input test", func(t *testing.T) {
		input, err := input.FileToLines("input.txt")
		require.NoError(t, err, "failed to read in input")
		result := part1(input)
		t.Log(result)
	})
}
func TestPart2(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		assert.Equal(t, "#3", part2([]string{
			"#1 @ 1,3: 4x4",
			"#2 @ 3,1: 4x4",
			"#3 @ 5,5: 2x2",
		}))
	})

	t.Run("input test", func(t *testing.T) {
		input, err := input.FileToLines("input.txt")
		require.NoError(t, err, "failed to read in input")
		result := part2(input)
		t.Log(result)
	})
}
