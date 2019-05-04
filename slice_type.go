package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// length
	fmt.Println(len(loons)) //3

	fmt.Println("----")
	// 0 indexing
	fmt.Println(loons[1]) //daffy

	fmt.Println("----")
	// slices
	fmt.Println(loons[1:]) // [daffy taz]

	// iterating over slice
	fmt.Println("----")
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// range keyword, kind of like Python foreach
	fmt.Println("----")
	for i := range loons {
		fmt.Println(i)
	}

	// Go loops always has a counter
	fmt.Println("----")
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	// if you don't want to capture it, use _
	// unused variables in Go are compilation errors
	fmt.Println("----")
	for _, name := range loons {
		fmt.Println(name)
	}

	// extend a slice with append
	fmt.Println("----")
	loons = append(loons, "elmer")
	fmt.Println(loons)
}
