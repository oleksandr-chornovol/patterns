package builder

import "fmt"

type ScheduledOrder struct {
	Items            []string
	Price            float64
	TotalBookingTime float64
}

func (o ScheduledOrder) Show() {
	fmt.Printf("Price: %f\n", o.Price)
	fmt.Printf("Booking time: %f\n", o.TotalBookingTime)
}
