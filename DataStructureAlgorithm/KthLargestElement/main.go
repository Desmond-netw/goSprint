package main

// func to take my input and out
func FindKthLargest(nums []int, k int) int {
	// step 1: determine our largest index (value)
	//thus len(nums) - k
	// so let call this targetPosition
	targetPosition := len(nums) - k

	// step 2: set up pointers for the search in array
	// left and right pointer.
	left, right := 0, len(nums)-1

	// step3  : let partition the slice
	for left <= right {

		// step4 : partition the array
		pivotIndex := partition(nums, left, right)
	}

}

func partition(nums []int, l, r int) int {

	pivot := nums[l]

	left, right := l+1, r

	for left <= right {

		for left <= right && nums[left] < pivot {
			left++
		}

		for left <= right && nums[right] > pivot {
			right--
		}

		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[right], nums[left] = nums[left], nums[right]

}
