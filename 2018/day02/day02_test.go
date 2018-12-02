package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
}
