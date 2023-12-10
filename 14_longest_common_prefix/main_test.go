package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name:     "fl",
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "no prefix",
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "one str",
			strs:     []string{"dog"},
			expected: "dog",
		},
		{
			name:     "empty strs",
			strs:     []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, longestCommonPrefix(tt.strs))
		})
	}
}
