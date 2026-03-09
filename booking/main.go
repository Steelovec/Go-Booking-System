package main

import (
	"fmt"
	"strings"
)

func main(){

	// alternative way of declaring vars, cant do const and the explicit types like uint
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	// this wont be ever negative bc of unit
	var remainingTickets uint = 50 

	// its fixed, u cant mix the types like ints or strings,...
	bookings := []string{}
	
	fmt.Printf("Welcome to %v booking app. \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	for {
		// gives the var a data type
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		// has to be a pointer
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)	

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets you want: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)	


		firstNames := []string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0]) 
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)


	}

}

// 1:24