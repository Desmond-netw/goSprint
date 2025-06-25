package main

import (
	"reflect"
	"testing"
)

func FilterBySumTest(t *testing.T) {
	// initilaize var test struct
	tests := []struct {
		name       string
		inputarray [][]int
		limit      int
		expected   [][]int
	}{
		{
			name: "Basic example",
			inputarray: [][]int{
				{1, 2, 3},
				{4, 5},
				{0, 1, 0, 1},
				{10},
			},
			limit: 5,
			expected: [][]int{
				{1, 2, 3},
				{4, 5},
				{10},
			},
		},
		{
			name: "Empty subarray",
			inputarray: [][]int{
				{},
				{5, 6},
			},
			limit: 1,
			expected: [][]int{
				{5, 6},
			},
		},
		{
			name: "subarrays with negative numbers",
			inputarray: [][]int{
				{-2, -3}, {4, -1}, {-10, 5, 10},
			},
			limit: 0,
			expected: [][]int{
				{4, -1},
				{-10, 5, 10},
			},
		},
	}

	// iterate for testing
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FilterBySum(test.inputarray, test.limit)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Expected %v but got %v ", test.expected, got)
			}
		})
	}
}
