package main

//import
import (
	"fmt"
	"sort"
)

// type
type TrieNode struct {
	children map[rune]*TrieNode
	suggestions []string
}

// main function
func ProductSearchTool( products []string , searchWorld string ) [][]string {
	
	// sort the products
	sort.String(products)

	// set root of products
	root := &TrieNode{children: make(map[rune]*TrieNode)}

	// Build the Trie 
	for _,product := range products {

		// current product setting 
		current := root
		for _, ch := range product {
			//condition to check ch in root
			if _,exists := current.children[ch]; !exists {
				current.children[ch] = &TrieNode{children: make(map[run]*TrieNode)}
			}
			current = current.children[ch]

			// add all suggestions 
			if len(current.suggestions) < 3 {
				current.suggestions = append(current.suggestions, product)
			}
			
		}
	}

	
}