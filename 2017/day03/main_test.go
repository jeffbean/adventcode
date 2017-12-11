package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	x, y := findCorrdinates(18)
	assert.Equal(t, x, -2.0)
	assert.Equal(t, y, 1.0)

	x, y = findCorrdinates(23)
	assert.Equal(t, x, 0.0)
	assert.Equal(t, y, -2.0)

}

func TestSliceAdding(t *testing.T) {
	x, y := part2BuildSpiral(10, 10)
	assert.Equal(t, 9, x)
	assert.Equal(t, 10, y)

	x, y = part2BuildSpiral(10, 25)
	assert.Equal(t, 11, x)
	assert.Equal(t, 9, y)

	x, y = part2BuildSpiral(10, 26)
	assert.Equal(t, 12, x)
	assert.Equal(t, 9, y)

	x, y = part2BuildSpiral(10, 806)
	assert.Equal(t, 10, x)
	assert.Equal(t, 8, y)

	total := part2(806)
	assert.Equal(t, 2.0, total)

	total = part2(147)
	assert.Equal(t, 4.0, total)

}

func TestSummingCell(t *testing.T) {
	testingBegin := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 4, 2, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	testingNext := [][]int{
		{0, 0, 0, 0, 0},
		{0, 5, 4, 2, 0},
		{0, 10, 1, 1, 0},
		{0, 11, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	tests := []struct {
		inGrid [][]int
		x, y   int
		want   int
	}{
		{
			inGrid: testingBegin,
			x:      1,
			y:      2,
			want:   4,
		},
		{
			inGrid: testingBegin,
			x:      1,
			y:      1,
			want:   5,
		},
		{
			inGrid: testingNext,
			x:      3,
			y:      2,
			want:   23,
		},
	}

	for _, tt := range tests {
		grid := &grid{
			m: tt.inGrid,
		}
		val := grid.getValue(tt.x, tt.y)
		assert.Equal(t, tt.want, val)

	}
}
