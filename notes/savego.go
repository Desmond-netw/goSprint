package main

import (
	"fmt"
	"os"
)

// Write notes back to file
func saveNotes(fileName string, notes []string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	
	for _, note := range notes {
		_, err := file.WriteString(note + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
