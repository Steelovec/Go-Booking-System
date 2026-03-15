package main

import (
	"fmt"
	"strings"
	"booking/helper"
)

	// alternative way of declaring vars, cant do const and the explicit types like uint
	var conferenceName = "Go conference"
	const conferenceTickets int = 50
	// this wont be ever negative bc of unit
	var remainingTickets uint = 50

	// its fixed, u cant mix the types like ints or strings,...
	var bookings = []string{}

func main() {


	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket( userTickets ,  firstName , lastName , email , )

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email doesnt contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app. \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")
}

// first we have parameters, then the output paremeters
func getFirstNames() []string {
	// this gets the first name from the slice
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}


func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket( userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}


// 2:33
