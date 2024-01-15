package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Ticket Swift"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	for {
		fmt.Print("Enter your first name = ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name = ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address = ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets = ")
		fmt.Scan(&userTickets)

		isvalidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets

		if isvalidName && isValidEmail && isValidTicket {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			//fmt.Println(len(bookings))

			firstNames := []string{}
			for _, booking := range bookings {
				nameSlice := strings.Fields(booking)
				firstNames = append(firstNames, nameSlice[0])
			}
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			fmt.Printf("This are all our bookings = %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("No tickets remaining.Please come back and try again later.")
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
