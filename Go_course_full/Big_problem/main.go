package main

import "fmt"

func main() {

	fmt.Println("\tFactorial using Iterative Approach")
	fmt.Println("\t", FactorialIterative(5))

	fmt.Println("\tFactorial with Recursive Approach")
	fmt.Println("\t", FactorialRecursive(0))

	fmt.Println("\tTo The Power function :")
	fmt.Println("\t", ToThePowerIterative(4, 2))
}
