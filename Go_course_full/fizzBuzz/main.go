package main

import "fmt"

func fizzBuzz(n int32) {
    for i := int32(1); i <= n; i++ {
        switch {
        case i%3 == 0 && i%5 == 0:
            fmt.Println("FizzBuzz")
        case i%3 == 0:
            fmt.Println("Fizz")
        case i%5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}

func main() {
    fizzBuzz(12) // Example: Prints FizzBuzz for 1 to 15
}