package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	test := []string{"aa bb cc dd ee", "aa bb cc dd aa"}
	count := validPassphrases(test, validationPart1)
	assert.Equal(t, 1, count)
}

func TestPart2(t *testing.T) {
	test := []string{"abcde fghij", "abcde xyz ecdab"}
	count := validPassphrases(test, validationPart2)
	assert.Equal(t, 1, count)

	test = []string{"abcde fghij", "abcde xyz ecdab", "iiii oiii ooii oooi oooo", "oiii ioii iioi iiio", "a ab abc abd abf abj"}
	count = validPassphrases(test, validationPart2)
	assert.Equal(t, 3, count)
}
