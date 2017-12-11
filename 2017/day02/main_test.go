package main

import (
	"testing"

	"github.com/jeffbean/adventcode/2017/common"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input [][]int
		want  int
	}{
		{
			input: [][]int{
				{5, 1, 9, 5},
				{7, 5, 3},
				{2, 4, 6, 8},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, part1(tt.input))
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input [][]int
		want  int
	}{
		{
			input: [][]int{
				{5, 9, 2, 8},
				{9, 4, 7, 3},
				{3, 8, 6, 5},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, part2(tt.input))
	}
}

func TestMatrix(t *testing.T) {
	data := common.ReadFile("fixtures/one.txt")
	ss := readInSpreadsheet(data)
	want := [][]int{
		{5, 9, 2, 8},
		{9, 4, 7, 3},
		{3, 8, 6, 5},
	}
	assert.Equal(t, want, ss)
}
