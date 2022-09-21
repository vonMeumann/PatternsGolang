package main

import "awesomeProject/patterns"

var (
	ticketPrice = 59.99
	strategies  = []patterns.Strategy{ // определяем массив стратегий
		&patterns.FirstTimePerform{},
		&patterns.CalendarPrice{},
		&patterns.NewAlbumPrice{},
	}
)

func main() {
	calculator := patterns.Calculator{} // объявление калькулятора цен
	for _, strategy := range strategies {
		calculator.SetStrategy(strategy)
		calculator.PriceCalculation(ticketPrice)
	}
}
