package main

import "fmt"

func main() {
	// test := []int{1, 2, 3, 4, 5, 6, 7}
	neg_test := []int{-7, -6, -5, -4, -1}

	var max int = 0
	for pos, value := range neg_test {
		if pos == 0 {
			max = value
		} else if value > max {
			max = value
		}
	}
	fmt.Printf("Max is %d\n", max)
}

// cleaner solution, imo
// func main() {
// 	neg_test := []int{-7, -6, -5, -4, -1}

// 	max := neg_test[0]
// 	for _, value := range neg_test[1:] {
// 		if value > max {
// 			max = value
// 		}
// 	}
// 	fmt.Println(max)
// }
