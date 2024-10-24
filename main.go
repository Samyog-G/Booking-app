package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50

	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	for remainingTickets > 0 {

		var firstName string
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		var lastName string
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		var email string
		fmt.Println("Enter your email")
		fmt.Scan(&email)

		var userTickets int
		fmt.Println("Enter number of tickets you want to buy")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		isValidEmail := strings.Contains(email, "@") //strings package has Contains which checks whether the given value is present or not
		isValidTicketNumber := userTickets < 0 && userTickets <= remainingTickets

		//if user wants to book more tickets than remaining tickets
		if isValidEmail && isValidTicketNumber {

			var bookings []string
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank You %v %v for booking %v tickets. You will receive the email of your tickets at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remain for the conference %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			//this lines below defines whole logic of getting first name of the user
			for _, booking := range bookings { //_ is the blank identifier (ignores the unused variable ) used when you want
				var names = strings.Fields(booking)       //splits the booking string into words based on whitespace
				firstNames = append(firstNames, names[0]) //append the first name to the list
				//}

				// fmt.Printf("First names:%v\n", firstNames) //print the first names
				fmt.Printf("These are all our bookings: %v\n", firstNames) //only provides first name of the people who bought tickets
			}

			if remainingTickets == 0 {
				fmt.Println("Our conference is fully booked. Come back next year.")
				break
			}

		} else if userTickets == remainingTickets {
			fmt.Println("Sorry all the tickets have been sold out")
		} else {
			if !isValidEmail {
				fmt.Printf("Invalid email.\n")
			}

			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you entered is invalid.\n")
			}
			fmt.Printf("Your input data is invalid\n\n")
			continue

		}

	}
}
