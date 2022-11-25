package main

import (
	"fmt"
	"patterns/adapter"
)

func main() {
	hole := adapter.NewHole(5)
	roundPeg := adapter.NewRoundPeg(5)
	fmt.Println(hole.Fit(roundPeg))

	smallSquarePeg := adapter.NewSquarePeg(5)
	largeSquarePeg := adapter.NewSquarePeg(10)
	//fmt.Println(hole.Fit(smallSquarePeg))

	smallSquarePegAdapter := adapter.NewSquarePegAdapter(smallSquarePeg)
	largeSquarePegAdapter := adapter.NewSquarePegAdapter(largeSquarePeg)
	fmt.Println(hole.Fit(smallSquarePegAdapter))
	fmt.Println(hole.Fit(largeSquarePegAdapter))
}
