package main

import "fmt"

// filter 2d array and print the limit
/*
----------------- Algorithm -------------------
1. Initilize the empty [][] array and limit
2. interate throught the  arr for subarray[]
	- initia value of sum := 0
3. for each subarry []
- compute is sum
- condition if sum >= limit , append to output
4. print out output
*/

func FilterBySum(arr [][]int, limit int) [][]int {
	// output
	output := [][]int{}

	// get sub array
	for _, subarray := range arr {

		// initilize sum
		sum := 0
		for _, value := range subarray {
			sum += value
		}
		// condition to find limt
		if sum >= limit {
			output = append(output, subarray)
		}

	}
	return output
}

// testing
func main() {
	arr := [][]int{
		{1, 3, 2},
		{4, 5},
		{0, 1, 0, 1},
		{2},
	}

	fmt.Println("\t\v\nFilter array :", FilterBySum(arr, 5))
}
