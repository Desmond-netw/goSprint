package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read all notes from file
func loadNotes(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()
	
	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		note := scanner.Text()
		if note != "" {
			notes = append(notes, note)
		}
	}
	
	return notes
}
