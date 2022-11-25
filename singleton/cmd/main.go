package main

import "patterns/singleton"

func main() {
	i1 := singleton.GetInstance()
	i2 := singleton.GetInstance()
	i3 := singleton.GetInstance()

	i1.DoSomething()
	i2.DoSomething()
	i3.DoSomething()
}
