package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	test := []string{"aa bb cc dd ee", "aa bb cc dd aa"}
	count := validPassphrases(test)
	assert.Equal(t, 1, count)
}
