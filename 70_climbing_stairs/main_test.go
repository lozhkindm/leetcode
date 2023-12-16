package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "1",
			n:        1,
			expected: 1,
		},
		{
			name:     "2",
			n:        2,
			expected: 2,
		},
		{
			name:     "3",
			n:        3,
			expected: 3,
		},
		{
			name:     "4",
			n:        4,
			expected: 5,
		},
		{
			name:     "5",
			n:        5,
			expected: 8,
		},
		{
			name:     "45",
			n:        45,
			expected: 1836311903,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, climbStairs(tt.n))
		})
	}
}
