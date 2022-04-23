package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const numAvalibleTickets = 50
	var avalibleTickets = numAvalibleTickets
	fmt.Printf("Welcome to our %v where you can buy tickets.\n", conferenceName)
	fmt.Printf("Currently there are %v Tickets left for sale out of %v\n", avalibleTickets, numAvalibleTickets)

}
