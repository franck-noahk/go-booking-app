package main

import (
	"booking-app/internal/booking"
	"fmt"
)

func main() {
	// int, int8, int16, int32, int64,
	// uInt ⟶ unsighened integer
	const numMaxTickets = 50

	var conferenceName = "Go Conference"
	var firstName string
	var userTickets uint64
	var avalibleTickets uint = numMaxTickets
	userName := "Hi"

	var bookings []booking.Buyer = []booking.Buyer{}

	fmt.Printf("Welcome to our %v where you can buy tickets.\n", conferenceName)
	fmt.Printf("Currently there are %v Tickets left for sale out of %v\n", avalibleTickets, numMaxTickets)

	avalibleTickets = booking.FillPeople(bookings, avalibleTickets)

	if userTickets > numMaxTickets {
		fmt.Printf("%v you are unable to book %v tickets.\nThere are only %v tickets avalible for purchase.\n", firstName, userTickets, avalibleTickets)
	} else {
		fmt.Printf("%v booked %v tickets\nThere are currently %v tickets avalible.\n", userName, userTickets, avalibleTickets)
	}
}

// func getAvalibleTickets() {
// 	fmt.Printf("Currently there are %v Tickets left for sale out of %v\n", avalibleTickets, numAvalibleTickets)
// }
