package helpers

import (
	"fmt"
	"net/mail"
	"sync"
	"time"
)

func ValidateUserDataInput(firstName string, lastName string, email string, availableTickets uint, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	_, err := mail.ParseAddress(email)
	isValidEmail := err == nil
	// ticketType := reflect.TypeOf(userTickets).Kind()
	// isValidTicket := ticketType == reflect.Uint && userTickets <= availableTickets
	isValidTicket := userTickets <= availableTickets
	return isValidName, isValidEmail, isValidTicket
}

func GreetUsers(confName string, numberOfTickets uint, availableTickets uint) {
	fmt.Printf("Hello every One and Welcome to %v conference!\n", confName)
	fmt.Printf("There are %v in total tickets and %v tickets available!\n", numberOfTickets, availableTickets)
}

func GetUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var tickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter your requested number of tickets: ")
	fmt.Scan(&tickets)

	return firstName, lastName, email, tickets
}

type User struct {
	FirstName  string
	LastName   string
	Email      string
	TicketsNum uint
}

var bookings = make([]User, 0)

func BookTicket(firstName string, lastName string, email string, numOfTickets uint) []User {
	user := User{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		TicketsNum: numOfTickets,
	}
	bookings = append(bookings, user)
	return bookings
}

func SendEmail(email string, userData User, wg *sync.WaitGroup) {
	firstName := userData.FirstName
	fmt.Printf("Sending Email to %v ", firstName)
	time.Sleep(5 * time.Second)
	wg.Done()
}
