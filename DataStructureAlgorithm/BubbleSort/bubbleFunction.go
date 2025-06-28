package main

// function to sort array using bubble sorting

func Bubble_Sort(numbs []int) {
	// step 2: find the len of array

	n := len(numbs) // to get the lenth to prevent exceding

	// step 3: iterate over array

	for i := 0; i < (n - 1); i++ {
		swapped := false
		for j := 0; j < (n - i - 1); j++ {
			if numbs[j] > numbs[j+1] {
				numbs[j], numbs[j+1] = numbs[j+1], numbs[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
