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
		
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			
			firstNames := getFirstNames(bookings)
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

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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
	
	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets -= userTickets
	// array: bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets lef for %v.\n", remainingTickets, conferenceName)
}
