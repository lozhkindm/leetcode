package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "()",
			s:        "()",
			expected: true,
		},
		{
			name:     "()[]{}",
			s:        "()[]{}",
			expected: true,
		},
		{
			name:     "(]",
			s:        "(]",
			expected: false,
		},
		{
			name:     "mod 2",
			s:        "[",
			expected: false,
		},
		{
			name:     "empty stack",
			s:        "]",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, isValid(tt.s))
		})
	}
}
