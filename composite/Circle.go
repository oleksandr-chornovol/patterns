package composite

import "fmt"

type Circle struct {
	radius int
	Dot
}

func NewCircle(x, y, radius int) *Circle {
	return &Circle{
		Dot: Dot{
			x: x,
			y: y,
		},
		radius: radius,
	}
}

func (c Circle) Draw() {
	fmt.Printf("Draw circle in %d, %d, with radius %d\n", c.x, c.y, c.radius)
}
