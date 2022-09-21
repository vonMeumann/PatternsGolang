package patterns

import "fmt"

type NewAlbumPrice struct {
}

func (c *NewAlbumPrice) PriceCalculation(ticketPrice float64) {
	discount := 20
	totalPrice := ticketPrice - (ticketPrice*float64(discount))/100
	fmt.Printf("Common price: [%f], as the band will perform their new album, your dicount is [%d], the total price for ticket is [%f]", ticketPrice, discount, totalPrice)
	fmt.Println("")
}
