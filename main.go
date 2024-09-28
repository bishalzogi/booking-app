package main

import "fmt"

func main() {
	var confrenceName = "Go confrence"
	const confrenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("confranceTickets is %T, remainingTickets is %T, confrenceName is %T\n", confrenceTickets, remainingTickets, confrenceName)

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v tickets available\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attende")

	var userName string
	var userTickets int
	// ask user for their name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
