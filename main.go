package main

import "fmt"

func main() {
	conferenceName := "Ticket Swift"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)
}
