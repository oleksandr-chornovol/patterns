package builder

import "math/rand"

type SimpleOrderBuilder struct {
	order SimpleOrder
}

func (sob SimpleOrderBuilder) SetItems(items []string) OrderBuilderInterface {
	sob.order.Items = items

	sob.order.Weight = rand.Float64() // mock logic of weight calculation from items
	sob.order.Price = rand.Float64()  // mock logic of price calculation from items

	return sob
}

func (sob SimpleOrderBuilder) SetDiscount(discount float64) OrderBuilderInterface {
	sob.order.Price -= discount

	return sob
}

func (sob SimpleOrderBuilder) GetOrder() OrderInterface {

	return sob.order
}
