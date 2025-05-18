package main

import (
	"bufio"
	"fmt"
	// "os"
	// "strings"
)

// Add a note to the collection
func addNote(fileName string, reader *bufio.Reader) {
	fmt.Println("\nEnter the note text:")
	noteText := getInput(reader)

	if noteText == "" {
		fmt.Println("Note cannot be empty.")
		return
	}

	notes := loadNotes(fileName)
	notes = append(notes, noteText)
	saveNotes(fileName, notes)
}
