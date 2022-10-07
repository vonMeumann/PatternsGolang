package main

import (
	"awesomeProject"
	"fmt"
)

var s, s1 string
var x, x1 int

func main() {

	Ticket := &awesomeProject.Ticket{} //Base
	fmt.Println("The Price of the Ticket is: ", Ticket.GetPrice())
	fmt.Println("Do you want to add something?")
	fmt.Scanln(&s)
	if s == "yes" {
		fmt.Println("What do you want to add?")
		fmt.Scanln(&s1)
		switch s1 {
		case "PlusMerch":
			Ticket2 := &awesomeProject.PlusMerch{TicketFrom: Ticket}
			fmt.Println("How many do you want?")
			fmt.Scanln(&x)
			for i := 1; i < x; i++ {
				Ticket2 = &awesomeProject.PlusMerch{TicketFrom: Ticket2}
			}
			fmt.Printf("The Price of the Ticket with  %d PlusMerch is: %d", x, Ticket2.GetPrice())
		case "FanZone":
			Ticket2 := &awesomeProject.FanZone{TicketFrom: Ticket}
			fmt.Println("How many do you want?")
			fmt.Scanln(&x)
			for i := 1; i < x; i++ {
				Ticket2 = &awesomeProject.FanZone{TicketFrom: Ticket2}
			}
			fmt.Printf("The Price of t	he Ticket with  %d Fanzone is: %d", x, Ticket2.GetPrice())
		case "Both":
			Ticket2 := &awesomeProject.FanZone{TicketFrom: Ticket}
			fmt.Println("How many Fanzone do you want?")
			fmt.Scanln(&x)
			fmt.Println("How many PlusMerch do you want?")
			fmt.Scanln(&x1)
			for i := 1; i < x; i++ {
				Ticket2 = &awesomeProject.FanZone{TicketFrom: Ticket2}
			}
			Ticket3 := &awesomeProject.PlusMerch{TicketFrom: Ticket2}
			for i := 1; i < x1; i++ {
				Ticket3 = &awesomeProject.PlusMerch{TicketFrom: Ticket3}
			}
			fmt.Printf("The Price of the Ticket with  %d Fanzone and %d PlusMerch is: %d", x, x1, Ticket3.GetPrice())
		}
	}
}
