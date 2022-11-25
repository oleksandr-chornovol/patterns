package composite

type CompoundGraphic struct {
	children []GraphicInterface
}

func (cg *CompoundGraphic) Add(graphic GraphicInterface) {
	cg.children = append(cg.children, graphic)
}

func (cg *CompoundGraphic) Move(x, y int) {
	for i := 0; i < len(cg.children); i++ {
		cg.children[i].Move(x, y)
	}
}

func (cg *CompoundGraphic) Draw() {
	for i := 0; i < len(cg.children); i++ {
		cg.children[i].Draw()
	}
}
