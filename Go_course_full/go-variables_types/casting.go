/*
Instructions
Write a function that accepts a float64 value, rounds it to the nearest integer, and then returns the result as an int.
*/

package main

import (
	"math"
)

func Castings(n float64) int {
	value := int(math.Round(n))

	return value
}