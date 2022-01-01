package main

import "fmt"

func main() {
	var confName = "Go Conference"
	const totalTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("\nWelcome to %v\n", confName)
	fmt.Printf("Total number of tickets are %v and remaining are %v\n\n", totalTickets, remainingTickets)

	var userName string
	var userTickets uint
	var userEmail string

	for {
		if remainingTickets == 0 {
			fmt.Println("\nWe apologize but we have no tickets remaining. Please come back next year! Have a nice day!")
			break
		}
		fmt.Printf("Please enter your full name: ")
		fmt.Scan(&userName)

		fmt.Printf("How many tickets you would like to book: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			fmt.Printf("Please enter your email: ")
			fmt.Scan(&userEmail)

			bookings = append(bookings, userName)
			fmt.Println("---------------------------------------------------------------------------------------")
			fmt.Printf("Thank you %v for booking %v tickets. You will recieve confirmation email on %v\n", userName, userTickets, userEmail)
			fmt.Println("---------------------------------------------------------------------------------------")

			fmt.Println("Total number of bookings yet :")
			for _, booking := range bookings {
				println(booking)
			}

			remainingTickets = remainingTickets - userTickets

			fmt.Printf("\nTotal tickets remaining to be booked are: %v\n", remainingTickets)
		} else {
			fmt.Printf("\nTotal number of remaining tickets are %v and you have entered %v\n", remainingTickets, userTickets)
			continue
		}
	}

}
