package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

// constructor equivalent in Go
// write a function and return pointer to object
// function name should start with New
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol can't be empty")
	}
	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}
	if price <= 0.0 {
		return nil, fmt.Errorf("price must be >= 0 (was %f)", price)
	}

	// this will allocate on the heap
	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil
}

// (t *Trade) is known as a receiver in Golang
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	t, err := NewTrade("MSFT", 10, 99.98, true)

	if err != nil {
		fmt.Printf("error: can't create trade - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())
}
