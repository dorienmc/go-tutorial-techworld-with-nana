package main

import (
	"fmt"
)

// Package level variables
// Defined at the top outside all functions
// Can be accessed inside any of the functions
// Can be accessed by all files which are inside the same package
// Cannot use := for package variables
var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main(){		
	greetUsers()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
		
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, firstName, lastName, email)
			
			firstNames := getFirstNames()
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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	
	// create a map for a user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
		
	bookings = append(bookings, userData)
	fmt.Print("List of bookings is %v\n", bookings)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets lef for %v.\n", remainingTickets, conferenceName)
}
