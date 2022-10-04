package main

import (
	"awesomeProject"
	"fmt"
)

var (
	Tiny   = &awesomeProject.Shop{}
	Medium = &awesomeProject.Market{IceCreams: 100, Chocolate: 30.0, Wrapper: Tiny}
	Large  = &awesomeProject.SuperGIGAMegamarket{IceCreams: 100, Chocolate: 30.0, Wafle: 20.0, Wrapper: Medium}
)

func main() {
	fmt.Println("Hello,If you need to buy ice cream, you can go to the shop, the market or the SuperGIGAMegamarket, but the price will be different in each place.")
	fmt.Println("Input, please, the name of Market where you want to buy an ice cream:")
	var name string
	fmt.Scan(&name)
	switch name {
	case "Shop":
		fmt.Println("The price of ice cream in the shop is", Tiny.GetPrice())
	case "Market":
		fmt.Println("The price of ice cream in the market is", Medium.GetPrice())
	case "SGM":

		fmt.Println("The price of ice cream in the SuperGIGAMegamarket is", Large.GetPrice())
	}

}
