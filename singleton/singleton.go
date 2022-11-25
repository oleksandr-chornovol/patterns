package singleton

import (
	"fmt"
	"sync"
)

//var once sync.Once

type single struct {
}

func (s single) DoSomething() {
	fmt.Println("DoSomething")
}

var singleInstance *single

func GetInstance() *single {
	var once sync.Once

	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
