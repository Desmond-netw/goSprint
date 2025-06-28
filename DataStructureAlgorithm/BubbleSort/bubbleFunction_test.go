package main

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {

	// internal struck for testcase
	testCase := []struct {
		name     string
		arrInput []int
		expected []int
	}{
		{
			name:     "Typical Reversal Unsorted List( westcase)",
			arrInput: []int{11, 12, 9, 7, 3},
			expected: []int{3, 7, 9, 11, 12},
		},
		{
			name:     "Already Sorted List",
			arrInput: []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse Sorted List (Worst Case)",
			arrInput: []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "List with Duplicate Elements",
			arrInput: []int{5, 2, 8, 2, 5},
			expected: []int{2, 2, 5, 5, 8},
		},
		{
			name:     "List with Negative Numbers",
			arrInput: []int{-5, 10, -1, 0, 3},
			expected: []int{-5, -1, 0, 3, 10},
		},
		{
			name:     "Empty List (Edge Case)",
			arrInput: []int{},
			expected: []int{},
		},
		{
			name:     "Single Element List (Edge Case)",
			arrInput: []int{42},
			expected: []int{42},
		},
	}

	// iterate of sort
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {

			// got
			got := make([]int, len(tt.arrInput))
			copy(got, tt.arrInput)

			Bubble_Sort(got)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Bubble Sort got %v, expected %v", got, tt.expected)
			}
		})
	}

}
