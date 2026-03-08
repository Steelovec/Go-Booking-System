package main

import(
	"fmt"
)

func main(){
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50 
		
	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var userName = ""
	// ask user for their name

	userName = "Tom"
	fmt.Println(userName)
}