package composite

import "fmt"

type Dot struct {
	x, y int
}

func NewDot(x, y int) *Dot {
	return &Dot{
		x: x,
		y: y,
	}
}

func (d *Dot) Move(x, y int) {
	d.x += x
	d.y += y
}

func (d *Dot) Draw() {
	fmt.Printf("Draw dot in %d, %d\n", d.x, d.y)
}
