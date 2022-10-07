package awesomeProject

type PlusMerch struct {
	TicketFrom Wrapper
}

func (c *PlusMerch) GetPrice() int {
	PlusMerchPrice := c.TicketFrom.GetPrice()
	return PlusMerchPrice + 7
}
