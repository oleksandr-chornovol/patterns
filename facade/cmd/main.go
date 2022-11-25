package main

import (
	"fmt"
	"patterns/facade"
)

func main() {
	sf := facade.SimpleFacade{}
	fmt.Println(sf.GetSomething())
}
