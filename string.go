package main

import (
	"fmt"
)

func main() {
	book := "Thee colour of magic"
	fmt.Println(book)

	// Get length of string
	fmt.Println(len(book))
	// You can access bytes of string just like in C!
	// Yup! Strings are immutable like in C
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// You can access part of strings with Slice
	// [number(inclusive):number(not inclusive)]
	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println(book[:4])
	// use + to concatenate strings
	fmt.Println("t" + book[1:])
	// unicode
	fmt.Println("it was ≈ç≈ç˜∂∆∑ø©˜®´µ˜ price!")
	//multiline
	poem := `
	The road goes ever on
	Down from the door where it began
	...
	`
	fmt.Println(poem)
}
