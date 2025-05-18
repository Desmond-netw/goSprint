package main

import (
	"bufio"
	"os"
	"strings"
)

// Get a buffered reader for user input
func getReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

// Get user input and trim whitespace
func getInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
