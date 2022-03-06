package main

import (
	"fmt"
	"pr-go/booking-app/helpers"
	"sync"
)

// Wait for the concurrent jobs to finish
var wg = sync.WaitGroup{}

func main() {
	coferenceName := "Go-PR"
	const numOfTickets = 50
	var availableTickets uint = 50
	for {
		if availableTickets == 0 {
			fmt.Println("No more tickets available, sorry and bye!")
			break
		}
		helpers.GreetUsers(coferenceName, numOfTickets, availableTickets)
		firstName, lastName, email, userTickets := helpers.GetUserInputs()
		isValidName, isValidEmail, isValidTickets := helpers.ValidateUserDataInput(firstName, lastName, email, availableTickets, userTickets)
		if isValidEmail && isValidName && isValidTickets {
			bookings := helpers.BookTicket(firstName, lastName, email, userTickets)
			var lastBooking helpers.User = bookings[len(bookings)-1]
			availableTickets -= userTickets
			fmt.Printf("Last user: %v booked %v tickets and %v tickets are available\n", lastBooking.FirstName, lastBooking.TicketsNum, availableTickets)
			fmt.Println("Sending tickets to the user via email...")
			wg.Add(1)
			go helpers.SendEmail(email, lastBooking, &wg)

		} else {
			if !isValidEmail {
				fmt.Println("Wrong email provided!")
			}
			if !isValidName {
				fmt.Println("Wrong name provided!")
			}
			if !isValidTickets {
				fmt.Printf("Only %v tickets are available!\n", availableTickets)
			}
		}

	}
	wg.Wait()
}
