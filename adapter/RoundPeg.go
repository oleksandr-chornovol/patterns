package adapter

type RoundPeg struct {
	radius float64
}

func NewRoundPeg(radius float64) *RoundPeg {
	return &RoundPeg{
		radius: radius,
	}
}

func (rp RoundPeg) GetRadius() float64 {
	return rp.radius
}
