package awesomeProject

type Market struct {
	IceCreams int
	Chocolate float64
	Wrapper   Wrapper
}

func (item *Market) GetPrice() float64 {
	return item.Wrapper.GetPrice() + float64(item.IceCreams) + item.Chocolate
}
