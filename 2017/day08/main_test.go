package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testString = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func TestPart1(t *testing.T) {

	assert.Equal(t, 15, handleIncDec("inc", 5, 10))
	assert.Equal(t, -10, handleIncDec("inc", 0, -10))
	assert.Equal(t, -10, handleIncDec("dec", 0, 10))
	assert.Equal(t, 10, handleIncDec("dec", 0, -10))

	part1([]byte(testString))

}

func TestPart2(t *testing.T) {

}
