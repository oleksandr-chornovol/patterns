package adapter

import "math"

type SquarePegAdapter struct {
	squarePeg *SquarePeg
}

func NewSquarePegAdapter(squarePeg *SquarePeg) *SquarePegAdapter {
	return &SquarePegAdapter{
		squarePeg: squarePeg,
	}
}

func (spa SquarePegAdapter) GetRadius() float64 {
	return spa.squarePeg.GetWidth() / math.Sqrt(2)
}
