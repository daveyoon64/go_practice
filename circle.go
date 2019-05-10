package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

type Circle struct {
	Center Point
	Radius int
}

// constructor
func NewCircle(x int, y int, radius int) (*Circle, error) {
	if x < 0 {
		return nil, fmt.Errorf("x must be greater than 0")
	}
	if y < 0 {
		return nil, fmt.Errorf("y must be greater than 0")
	}
	if radius < 0 {
		return nil, fmt.Errorf("length must be > 0")
	}
	c := &Circle{
		Center: Point{x, y},
		Radius: radius,
	}
	return c, nil
}

func main() {
	c, err := NewCircle(1, 1, 5)
	if err != nil {
		log.Fatalf("Error: Your circle is kaput!")
	}
	fmt.Printf("%+v\n", c)
}
