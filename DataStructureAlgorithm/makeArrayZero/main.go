package main

import "fmt"

func main() {
	// Application to find the Minimum number of operations
	// this method looks for all unique values
	//leetcode : 2357 Make array Zero by substracting Equal Amounts
	fmt.Println("-- Solutions 1 : ")
	fmt.Println("This method find all the unique non zero \nNumbers of the array and store it in a hashmap then print it len")
	fmt.Println("Input: {1, 5, 0, 3, 5} ")
	fmt.Println("OutPut: ", minimumOperation([]int{2, 5, 6, 0, 9, 5}))
}
