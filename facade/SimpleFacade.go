package facade

type SimpleFacade struct {
	cc1 ComplicatedClass1
	cc2 ComplicatedClass2
}

func (sf *SimpleFacade) GetSomething() int {
	return sf.cc1.someFunction() + sf.cc2.someFunction()
}
