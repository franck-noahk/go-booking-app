package booking

import (
	"booking-app/cmd"
	"fmt"
	"strconv"
)

func FillPeople(list []Buyer, ticketsAvalible uint) uint {
	for ticketsAvalible > 0 {
		firstName := cmd.PromptString("What is your first Name?")
		lastName := cmd.PromptString("What is your last name?")
		ticketsRequested, err := strconv.Atoi(cmd.PromptString("How many tickets do you want?"))
		if err != nil {
			fmt.Println("Must enter positive integers for tickets")
			continue
		}
		if firstName == "-1" || lastName == "-1" || ticketsRequested == -1 {
			break
		}

		if ticketsRequested > 0 && ticketsRequested <= int(ticketsAvalible) {
			ticketsAvalible = uint(ticketsAvalible) - uint(ticketsRequested)
			b := Buyer{firstName, lastName, ticketsRequested}
			list = append(list, b)
		}
		fmt.Printf("Currently there are %v.\n", ticketsAvalible)
	}
	return ticketsAvalible

}
