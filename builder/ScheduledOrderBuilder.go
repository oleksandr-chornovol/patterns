package builder

import "math/rand"

type ScheduledOrderBuilder struct {
	order ScheduledOrder
}

func (sob ScheduledOrderBuilder) SetItems(items []string) OrderBuilderInterface {
	sob.order.Items = items

	sob.order.TotalBookingTime = rand.Float64() // mock logic of calculation time from items

	return sob
}

func (sob ScheduledOrderBuilder) SetDiscount(discount float64) OrderBuilderInterface {
	sob.order.Price -= discount

	return sob
}

func (sob ScheduledOrderBuilder) GetOrder() OrderInterface {

	return sob.order
}
