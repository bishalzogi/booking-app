package main

import "fmt"

func main() {
	var confrenceName string = "Go confrence"
	const confrenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v tickets available\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attende")

	for {
		var firstName string
		var lastName string
		var userEmail string
		var userTickets uint
		// ask user for their name

		fmt.Print("Enter your First Name: \n")
		fmt.Scan(&firstName)

		fmt.Print("Enter your Last Name: \n")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: \n")
		fmt.Scan(&userEmail)

		fmt.Print("Enter Number of Tickets: \n")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you for booking the tickets %v %v\n ", firstName, lastName)

		fmt.Printf("%v Tickts are booked for you and you will soon recive conformation email in following email address %v \n", userTickets, userEmail)

		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, confrenceName)

		fmt.Printf("These are all our bookings: %v \n", bookings)

	}

}
