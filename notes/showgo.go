package main

import (
	// "bufio"
	"fmt"
	// "os"
)

// Show all notes in the collection
func showNotes(fileName string) {
	notes := loadNotes(fileName)

	fmt.Println("\nNotes:")
	if len(notes) == 0 {
		fmt.Println("No notes in this collection.")
		return
	}

	for i, note := range notes {
		fmt.Printf("%03d - %s\n", i+1, note)
	}
}
