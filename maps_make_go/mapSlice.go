package main

import "fmt"

// now I get variable subjects map key : value where key name(string) value = slice[subjects]
//step 1: inilize my function with subjects make(map[string][]string)
// step2 : iterate over subjects and print name (key) value (subject)

func GetSubjects(name string, subjects []string) {
	// get input in finction parameters
	fmt.Println("\nApplication Print MapSlice")

	subjectsMap := make(map[string][]string)

	for key, subject := range subjectsMap {
		name = key
		subjects = subject
	}
	fmt.Println(name, subjects)
}
