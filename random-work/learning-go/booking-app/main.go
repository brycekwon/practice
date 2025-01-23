package main

import (
	"fmt"
	"strings"
)

const conferenceName string = "Go Conference"
const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []string

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

			fmt.Printf("These are all our bookings: %v.\n", firstNames(bookings))

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is to short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain '@'")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func firstNames(bookings []string) []string {
	firstNames := []string{}
	for _, e := range bookings {
		var names = strings.Fields(e)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
