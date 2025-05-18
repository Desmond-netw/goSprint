package main

import (
	"bufio"
	"fmt"
	"strconv"
)

// Delete a note from the collection
func deleteNote(fileName string, reader *bufio.Reader) {
	notes := loadNotes(fileName)
	
	if len(notes) == 0 {
		fmt.Println("\nNo notes to delete.")
		return
	}
	
	fmt.Println("\nEnter the number of note to remove or 0 to cancel:")
	input := getInput(reader)
	
	noteNum, err := strconv.Atoi(input)
	if err != nil || noteNum < 0 || noteNum > len(notes) {
		fmt.Println("Invalid note number.")
		return
	}
	
	if noteNum == 0 {
		return // User canceled the operation
	}
	
	// Remove the selected note
	notes = append(notes[:noteNum-1], notes[noteNum:]...)
	
	// Write back to file
	saveNotes(fileName, notes)
}
