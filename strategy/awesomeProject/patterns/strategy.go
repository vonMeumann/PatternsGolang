package patterns

type Strategy interface {
	PriceCalculation(ticketPrice float64)
}
