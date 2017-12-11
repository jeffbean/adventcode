package main

import (
	"testing"

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
