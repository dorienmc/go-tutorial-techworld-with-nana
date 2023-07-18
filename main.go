package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	
	// Array vs. slice (dynamic array)
	// array: var bookings [50]string //[50]string{}
	// slice
	var bookings []string
		
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for name
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		// ask for amount of tickets
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		
		// Input validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets
			// array: bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName + " " + lastName)
			
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("There are %v tickets lef for %v.\n", remainingTickets, conferenceName)
			
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
		} else {
			fmt.Println("Invalid user input, try again")
			if !isValidName {
				fmt.Println("Firstname or lastname is to short")
			}
			if !isValidEmail {
				fmt.Println("Email should contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
			// fmt.Printf("There are only %v remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
		}
	}
	
	fmt.Println("Our conference is booked out. Come back next year.")
}
