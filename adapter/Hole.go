package adapter

type Hole struct {
	radius float64
}

func NewHole(radius float64) *Hole {
	return &Hole{
		radius: radius,
	}
}

func (h *Hole) GetRadius() float64 {
	return h.radius
}

func (h *Hole) Fit(peg PegInterface) bool {
	return h.GetRadius() >= peg.GetRadius()
}
