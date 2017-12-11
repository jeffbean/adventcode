package common

import (
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
			sample: "abcde",
			others: []string{
				"abced",
				"xyz",
				"ecdab",
			},
			want: true,
		},
		{
			sample: "asdfasdfsadf",

			others: []string{
				"abcde",
				"xyz",
				"ecdab",
			},
			want: false,
		},
		{
			sample: "oiii",
			others: []string{"ioii"},
			want:   true,
		},
		{
			sample: "xiii",
			others: []string{"ixii", "iixi"},
			want:   true,
		},
		{
			sample: "biii",
			others: []string{"iiii", "bbii", "bbbi"},
			want:   false,
		},
		{
			sample: "abj",
			others: []string{"a", "ab", "abc", "abd", "abf"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.sample, func(t *testing.T) {
			isA := DetectAnagram(tt.sample, tt.others)
			assert.Equal(t, tt.want, isA)
		})
	}
}
