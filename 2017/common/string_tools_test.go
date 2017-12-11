package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	tests := []struct {
		sample string
		others []string
		want   bool
	}{
		{
			others: []string{
				"abced",
				"xyz",
				"ecdab",
			},
			want: true,
		},
		{
			others: []string{"ioii", "oiii"},
			want:   true,
		},
		{
			others: []string{"ixii", "iixi"},
			want:   true,
		},
		{
			others: []string{"iiii", "bbii", "bbbi"},
			want:   false,
		},
		{
			others: []string{"a", "ab", "abc", "abd", "abf"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.others), func(t *testing.T) {
			isA := DetectAnagram(tt.others)
			assert.Equal(t, tt.want, isA)
		})
	}
}
