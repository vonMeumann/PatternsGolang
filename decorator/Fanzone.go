package awesomeProject

type FanZone struct {
	TicketFrom Wrapper
}

func (c *FanZone) GetPrice() int {
	FanZonePrice := c.TicketFrom.GetPrice()
	return FanZonePrice + 10
}
