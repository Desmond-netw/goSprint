package main

import (
	"fmt"
	"strings"
)

// Application count the appearance of words
// input : "go is fun and go is fast "
// exptect output:  map[go:2 is:2 fun: 1 and:1 fast : 1]
func WordCount(text string) {
	// step1: get input
	fmt.Println("--- Word Apperance Counter - ")
	//step 2 : split inputs
	words := strings.Fields(text) // [ split of words taking space out]

	// step3: store words count
	counts := make(map[string]int)

	// when ever i iterate over words, i will store them in counts
	for _, word := range words {
		counts[word]++
	}
	fmt.Println(counts)

}
