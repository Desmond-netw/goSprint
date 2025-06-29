package main

// function to sort array in smallest to biggest

func SelectSorting(myArry []int) {

	// Equivalent of my_array:

	// find length of array
	n := len(myArry)

	// outer iterationg from second to last
	for i := 0; i < n-1; i++ {
		// asume i is minIndex
		minIndex := i

		for j := i + 1; j < n; j++ {
			if myArry[j] < myArry[minIndex] {
				minIndex = j
			}
			myArry[i], myArry[minIndex] = myArry[minIndex], myArry[i]
		}
	}
}
