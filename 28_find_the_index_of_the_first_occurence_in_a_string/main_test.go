package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "sadbutsad",
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			name:     "leetcode",
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			name:     "odyssey",
			haystack: "odyssey",
			needle:   "sey",
			expected: 4,
		},
		{
			name:     "odyssee",
			haystack: "odyssee",
			needle:   "sey",
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, strStr(tt.haystack, tt.needle))
		})
	}
}
