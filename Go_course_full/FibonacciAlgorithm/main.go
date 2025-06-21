package main

// Creating a fibonacci of 20
/*
Algorithm

1. Get variable for first two
2. print out first two numbers
3. Add the two numbers to get new number
4. update newWith with prveious number
5. print new number
*/
import "fmt"

func main() {
	// step 1
	prev2 := 0
	prev1 := 1

	fmt.Println(prev2)
	fmt.Println(prev1)

	for i := 0; i < 18; i++ {
		//step 3:
		newFibo := prev2 + prev1
		fmt.Println(newFibo)
		prev2 = prev1
		prev1 = newFibo
	}
}
