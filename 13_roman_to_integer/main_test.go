package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		roman    string
		expected int
	}{
		{
			name:     "III",
			roman:    "III",
			expected: 3,
		},
		{
			name:     "LVIII",
			roman:    "LVIII",
			expected: 58,
		},
		{
			name:     "MCMXCIV",
			roman:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, romanToInt(tt.roman))
		})
	}
}
