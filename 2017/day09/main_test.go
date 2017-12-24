package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 6, part1([]byte("{{{}}}")))
	assert.Equal(t, 16, part1([]byte("{{{},{},{{}}}}")))
	assert.Equal(t, 1, part1([]byte("{<a>,<a>,<a>,<a>}")))
	assert.Equal(t, 9, part1([]byte("{{<ab>},{<ab>},{<ab>},{<ab>}}")))

	assert.Equal(t, 9, part1([]byte("{{<!!>},{<!!>},{<!!>},{<!!>}}")))
	assert.Equal(t, 3, part1([]byte("{{<a!>},{<a!>},{<a!>},{<ab>}}")), "{{<a!>},{<a!>},{<a!>},{<ab>}}")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2([]byte("{{{}}}")))
	assert.Equal(t, 0, part2([]byte("{{{},{},{{}}}}")))
	assert.Equal(t, 4, part2([]byte("{<a>,<a>,<a>,<a>}")))
	assert.Equal(t, 8, part2([]byte("{{<ab>},{<ab>},{<ab>},{<ab>}}")))

	assert.Equal(t, 0, part2([]byte("{{<!!>},{<!!>},{<!!>},{<!!>}}")))
	assert.Equal(t, 17, part2([]byte("{{<a!>},{<a!>},{<a!>},{<ab>}}")), "{{<a!>},{<a!>},{<a!>},{<ab>}}")
	assert.Equal(t, 0, part2([]byte("<>")))
	assert.Equal(t, 10, part2([]byte("<{o\"i!a,<{i<a>")))
	assert.Equal(t, 2, part2([]byte("<{!>}>")))

}
