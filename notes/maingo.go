package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for correct number of arguments
	if len(os.Args) != 2 {
		printHelp()
		return
	}

	// Get the collection name from arguments
	collection := os.Args[1]

	// Check if help was requested
	if collection == "help" {
		printHelp()
		return
	}

	// Create or load the collection
	fileName := "database/" + collection + ".txt"
	ensureFileExists(fileName)

	// Main program loop
	fmt.Println("Welcome to the notes tool!")
	
	reader := getReader()

	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		choice := getInput(reader)

		switch choice {
		case "1":
			showNotes(fileName)
		case "2":
			addNote(fileName, reader)
		case "3":
			deleteNote(fileName, reader)
		case "4":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Print help message
func printHelp() {
	fmt.Println("Usage: ./notestool [TAG]")
}

// Ensure the database directory exists
func ensureFileExists(fileName string) {
	// Ensure the database directory exists
	os.MkdirAll("database", 0755)
	
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		file.Close()
	}
}
