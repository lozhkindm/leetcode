package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "beginning duplicates",
			nums:     []int{1, 1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "many duplicates",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "ending duplicates",
			nums:     []int{0, 1, 1},
			expected: []int{0, 1},
		},
		{
			name:     "no duplicates",
			nums:     []int{0, 1, 2},
			expected: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates(tt.nums)
			require.Equal(t, tt.expected, tt.nums[:k])
		})
	}
}
