package patterns

import "fmt"

type CalendarPrice struct {
}

func (c *CalendarPrice) PriceCalculation(ticketPrice float64) {
	discount := 10
	totalPrice := ticketPrice - (ticketPrice*float64(discount))/100
	fmt.Printf("Common price: [%f], but for these dates discount is [%d]. So the total price for ticket is [%f]", ticketPrice, discount, totalPrice)
	fmt.Println("")
}
