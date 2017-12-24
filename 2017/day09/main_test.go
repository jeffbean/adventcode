package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	// assert.Equal(t, 6, part1([]byte("{{{}}}")))
	// assert.Equal(t, 16, part1([]byte("{{{},{},{{}}}}")))
	// assert.Equal(t, 1, part1([]byte("{<a>,<a>,<a>,<a>}")))
	// assert.Equal(t, 9, part1([]byte("{{<ab>},{<ab>},{<ab>},{<ab>}}")))

	// assert.Equal(t, 9, part1([]byte("{{<!!>},{<!!>},{<!!>},{<!!>}}")))
	assert.Equal(t, 3, part1([]byte("{{<a!>},{<a!>},{<a!>},{<ab>}}")), "{{<a!>},{<a!>},{<a!>},{<ab>}}")
}

func TestPart2(t *testing.T) {

}
