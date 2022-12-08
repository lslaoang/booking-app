package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets= 50
	fmt.Printf("Welcome to %s booking application! \n", conferenceName)
	fmt.Printf("We  have total of %d tickets and %d are still available. \n", conferenceTickets, remainingTickets	)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets int
	//ask user to enter name

	userName = "Integer"
	userTickets = 3

	fmt.Printf("User %v booked %v tickets. \n",userName , userTickets)
	fmt.Printf("conferenceTickets is %T	conferenceName is %T remainingTickets is %T	userName is %T 	userTickets is %T",
	conferenceTickets, conferenceName, remainingTickets, userName, userTickets)

}