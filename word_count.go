package main

import (
	"fmt"
	"strings"
)

func main() {
	counts := map[string]int{}
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	words := strings.Fields(text)

	for _, word := range words {
		// this solution produces as error like this:
		// ./word_count.go:20:10: cannot use word (type int) as type string in map index
		// not sure why, since it should match the types?

		// if word, ok := counts[word]; ok {
		// 	counts[strings.ToLower(word)]++
		// } else {
		// 	counts[strings.ToLower(word)] = 1
		// }
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)
}
