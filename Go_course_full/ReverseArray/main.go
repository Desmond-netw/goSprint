// A function that reverse an array

package main

import "fmt"

// funct
func ReverseArray(arr []int) {
	// init start variable and end variable
	start, end := 0, len(arr)-1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func main() {
	myarr := []int{0, 2, 3, 5, 6, 39}
	fmt.Println("My origiginal array: ", myarr)

	ReverseArray(myarr)
	fmt.Println("Reverse array: ", myarr)
}
