package main

import "fmt"

func main() {
	var conferenceName = "Not a Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v Booking Application!!\n", conferenceName)
	fmt.Printf("Out of %v tickets, we have %v left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Click here to get your tickets.")

}
