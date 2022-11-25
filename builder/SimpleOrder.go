package builder

import "fmt"

type SimpleOrder struct {
	Items  []string
	Price  float64
	Weight float64
}

func (o SimpleOrder) Show() {
	fmt.Printf("Price: %f\n", o.Price)
	fmt.Printf("Weight: %f\n", o.Weight)
}
