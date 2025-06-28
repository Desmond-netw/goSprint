package main

// boble sorting algorithm

/*
1. Start get array []
2. Visit each value of array and compare with next values
3. if value is grate than next one, swap values so that highers comes first
4. Go through the array as manytimme as neede to finnish
*/

import "fmt"

func main() {
	myarray := []int{64, 34, 25, 12, 22, 11, 90, 5}
	fmt.Println("Original array: ", myarray)

	// consider array lenth
	n := len(myarray)

	// bobble sort
	for i := 0; i < (n - 1); i++ {
		for j := 0; j < (n - i - 1); j++ {
			if myarray[j] > myarray[j+1] {
				//swap elements
				myarray[j], myarray[j+1] = myarray[j+1], myarray[j]
			}
		}
	}

	fmt.Println("Sorted Bobble: ", myarray)
}
