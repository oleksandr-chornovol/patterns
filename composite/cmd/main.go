package main

import "patterns/composite"

func main() {
	tree1 := &composite.CompoundGraphic{}
	tree2 := &composite.CompoundGraphic{}

	tree1.Add(composite.NewDot(1, 1))
	tree1.Add(tree2)

	tree2.Add(composite.NewDot(11, 11))
	tree2.Add(composite.NewCircle(11, 11, 3))

	tree1.Draw()
}
