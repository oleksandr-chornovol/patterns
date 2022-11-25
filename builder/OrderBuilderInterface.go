package builder

type OrderBuilderInterface interface {
	SetItems(items []string) OrderBuilderInterface
	SetDiscount(discount float64) OrderBuilderInterface
	GetOrder() OrderInterface
}
