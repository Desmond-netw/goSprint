package main

// using hash methods

import (
	"fmt"
	"sort"
)

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

func main() {
	myarr := []int{2, 3, 1, 2, 4, 5}
	fmt.Println(hasDuplicateSorted(myarr))
}
