package main

import "fmt"

// print the applications

func main() {
	fmt.Println("\n--- Bubble Sort Function --")
	myarry := []int{7, 3, 9, 12, 11}
	Bubble_Sort(myarry)

	fmt.Println("Sorted Array\n", myarry)
}
