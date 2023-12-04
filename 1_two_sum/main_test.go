package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Case 1",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Case 3",
			nums:     []int{3, 3, 3, 3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "Case 4",
			nums:     []int{},
			target:   1,
			expected: []int{},
		},
		{
			name:     "Case 5",
			nums:     []int{1, 2, 3, 4, 5},
			target:   10,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, twoSum(tt.nums, tt.target))
		})
	}
}
