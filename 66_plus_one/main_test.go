package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{
			name:     "increment last",
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "increment last 2",
			digits:   []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			name:     "one digit",
			digits:   []int{9},
			expected: []int{1, 0},
		},
		{
			name:     "O(n)",
			digits:   []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, plusOne(tt.digits))
		})
	}
}
