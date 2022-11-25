package builder

type OrderDirector struct {
}

func (od OrderDirector) BuildOrder(builder OrderBuilderInterface) OrderInterface {
	items := []string{"one", "two", "three"}

	return builder.
		SetItems(items).
		SetDiscount(0.5).
		GetOrder()
}
