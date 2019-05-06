package main

import (
	"fmt"
)

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	// you can make sure a resource is cleaned up
	// with defer
	defer cleanup("A")
	defer cleanup("B")
	fmt.Println("worker")
	// defer is used to simplify functions that perform
	// various clean-up actions
}

func main() {
	worker()
}
