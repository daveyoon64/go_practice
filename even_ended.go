package main

import "fmt"

func main() {
	// find all even-ended numbers (first and last digt are the same)
	// from 1000 - 9999
	cnt := 0
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			n := i * j
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				cnt++
			}
		}
	}
	fmt.Printf("Total even-ended numbers is: %d\n", cnt)
}
