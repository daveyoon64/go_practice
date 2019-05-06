package main

import (
	"fmt"
)

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	// deference pointer and multiply it by 2
	*n *= 2
}

func main() {
	// golang passes maps and slices by reference
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	// golang passes primitives by value
	val := 10
	double(val)
	fmt.Println(val)

	// for primitives, do it with pointers!
	// & to send the address like in C!
	doublePtr(&val)
	fmt.Println(val)
}
