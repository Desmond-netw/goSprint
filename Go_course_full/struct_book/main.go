package main

import "fmt"

// initilize a strutur
type Book struct {
	title  string
	author string
	pages  int
}

func main() {
	library := Book{
		title:  "Home Made Food",
		author: "Linda Kia",
		pages:  25,
	}

	fmt.Println("Title: ", library.title)
	fmt.Println("Author:", library.author)
	fmt.Println("Pages: ", library.pages)
}
