package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{
			name:     "4",
			x:        4,
			expected: 2,
		},
		{
			name:     "8",
			x:        8,
			expected: 2,
		},
		{
			name:     "99",
			x:        99,
			expected: 9,
		},
		{
			name:     "101",
			x:        101,
			expected: 10,
		},
		{
			name:     "5",
			x:        5,
			expected: 2,
		},
		{
			name:     "8192",
			x:        8192,
			expected: 90,
		},
		{
			name:     "2147395599",
			x:        2147395599,
			expected: 46339,
		},
		{
			name:     "0",
			x:        0,
			expected: 0,
		},
		{
			name:     "1",
			x:        1,
			expected: 1,
		},
		{
			name:     "2",
			x:        2,
			expected: 1,
		},
		{
			name:     "3",
			x:        3,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, mySqrt(tt.x))
		})
	}
}
