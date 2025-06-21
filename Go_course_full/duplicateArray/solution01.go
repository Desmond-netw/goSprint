package main

// using hash methods

import "sort"

func hasDuplicateSorted(arr []int) bool {
	sort.Ints(arr)

	// initilize n
	n := len(arr)

	// iterate
	for i := 0; i < n-1; i++ {
		if arr[i] == arr[i+1] {
			return true
		}
	}
	return false

}
