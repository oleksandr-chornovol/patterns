package main

import "patterns/builder"

func main() {
	director := builder.OrderDirector{}

	simpleOrder := director.BuildOrder(builder.SimpleOrderBuilder{})
	scheduledOrder := director.BuildOrder(builder.ScheduledOrderBuilder{})

	simpleOrder.Show()
	scheduledOrder.Show()
}
