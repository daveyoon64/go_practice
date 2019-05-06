package main

import (
	"fmt"
)

// point is a 2d point
type Point struct {
	X int
	Y int
}

// move moves the point
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := &Point{1, 2}
	p.Move(2, 3)
	fmt.Printf("%+v\n", p)
}
