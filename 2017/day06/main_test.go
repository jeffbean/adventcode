package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	testBlocks := blocks{
		banks: []int{0, 2, 7, 0},
	}
	largetsIndex := findLargest(testBlocks)
	assert.Equal(t, 2, largetsIndex)

	testBlocks = blocks{
		banks: []int{0, 9, 7, 0},
	}
	assert.Equal(t, 1, findLargest(testBlocks))

	testBlocks = blocks{
		banks: []int{0, 2, 7, 0},
	}
	assert.Equal(t, 5, redistibuteBlocks(testBlocks))
}

func TestPart2(t *testing.T) {

}
