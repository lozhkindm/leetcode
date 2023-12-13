package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{
			name: "not empty lists",
			list1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			list2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name:     "empty lists",
			list1:    nil,
			list2:    nil,
			expected: nil,
		},
		{
			name:  "one empty list",
			list1: nil,
			list2: &ListNode{
				Val:  0,
				Next: nil,
			},
			expected: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, mergeTwoLists(tt.list1, tt.list2))
		})
	}
}
