package patterns

import "fmt"

type FirstTimePerform struct {
}

func (c *FirstTimePerform) PriceCalculation(ticketPrice float64) {
	discount := 15
	totalPrice := ticketPrice - (ticketPrice*float64(discount))/100
	fmt.Printf("Common price: [%f], but for the first band's perform the discount is [%d], the total price for ticket is [%f]", ticketPrice, discount, totalPrice)
	fmt.Println("")
}
