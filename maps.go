package main

import (
	"fmt"
)

func main() {
	// keys and values have to be the same type
	stocks := map[string]float64{
		"AMzN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}

	// length of map
	fmt.Println(len(stocks))
	// get value of MSFT
	fmt.Println(stocks["MSFT"])
	// if not found, returns 0
	fmt.Println(stocks["TSLA"])
	// how to check if a key is in map
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}
	// setting a map
	stocks["TSLA"] = 322.12

	// delete a value
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// single value "for" is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	// double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
