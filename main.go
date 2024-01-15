package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName = "Ticket Swift"
var RemainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserData()

		isvalidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, RemainingTickets)

		if isvalidName && isValidEmail && isValidTicket {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames(bookings)
			fmt.Printf("This are all our bookings = %v\n", firstNames)

			if RemainingTickets == 0 {
				fmt.Println("No tickets remaining. Please come back and try again later.")
				break
			}
		} else {
			if !isvalidName {
				fmt.Println("First name or last name is too short. Please update")
			}
			if !isValidEmail {
				fmt.Println("Email address is not valid. Please update")
			}
			if !isValidTicket {
				fmt.Println("Number of tickets is not valid. Please update")
			}
		}

	}
}

func greetUsers() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", RemainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		nameSlice := strings.Fields(booking)
		firstNames = append(firstNames, nameSlice[0])
	}
	return firstNames

}

func getUserData() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Print("Enter your first name = ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name = ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address = ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets = ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName, lastName, email string) (uint, []string) {
	RemainingTickets = RemainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", RemainingTickets, conferenceName)
	return RemainingTickets, bookings
}

/*
Greet the User --> greetUsers()

Get the user input --> getUserData()

Validate the user input --> validateUserInput()

If everything is valid..

Book the tickets --> bookTickets()

Get First Names --> getFirstNames()

Print the first Names

If no tickets left , exit the program
*/
