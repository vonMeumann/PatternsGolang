package awesomeProject

type SuperGIGAMegamarket struct {
	IceCreams int
	Chocolate float64
	Wafle     float64
	Wrapper   Wrapper
}

func (item *SuperGIGAMegamarket) GetPrice() float64 {
	return item.Wrapper.GetPrice() + float64(item.IceCreams) + item.Chocolate + item.Wafle
}
