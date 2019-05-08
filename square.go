package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Center Point,
	Length int,
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

// move square
func (s *Square) Move(dx int, dy int) {
	
}

func (t *Square) Area() int {

}

// constructor
func NewSquare(x int, y int, length int) (*Square, error) {
	if x < 0 {
		return nil, fmt.Errorf("x must be greater than 0")
	}
	if y < 0 {
		return nil, fmt.Errorf("y must be greater than 0")
	}
	if length < 0 {
		return nil, fmt.Errorf("length must be > 0")
	}
	s := &Square{
		Center: Point{x, y},
		Length: length,
	}
	return s, nil
}