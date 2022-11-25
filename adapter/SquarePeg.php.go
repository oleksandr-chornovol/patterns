package adapter

type SquarePeg struct {
	width float64
}

func NewSquarePeg(width float64) *SquarePeg {
	return &SquarePeg{
		width: width,
	}
}

func (sp SquarePeg) GetWidth() float64 {
	return sp.width
}
