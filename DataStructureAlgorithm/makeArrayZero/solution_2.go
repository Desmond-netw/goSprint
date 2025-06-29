package main

import "math"

//step: find the non Negative integer array nums []int
//step: let x be smallest non zero number in the numbs of array
//step: let substract x from nums[]int elements
//step: print the number of times this was done,

func minimum_operation(nums []int) int {
	// set operations to 0
	operations := 0

	//find the smallest non zero number
	for {
		smallestNonZero := math.MaxInt32
		isNonZero := false

		for _, num := range nums {
			if num > 0 && num < smallestNonZero {
				smallestNonZero = num
				isNonZero = true
			}
		}
		// if no nonzero left , exist the loop
		if !isNonZero {
			break
		}

		// perform substractions
		for i := range nums {
			if nums[i] > 0 {
				nums[i] -= smallestNonZero
			}
		}

		operations++
	}

	return operations
}
