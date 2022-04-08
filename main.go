package main

import (
	"fmt"
	"strings"
)

const foodCompany string = "Shri Annapoorneshwari Bhojana"
const totalCoupons uint = 200

var remainingCoupons uint = 200
var bookings []string

func main() {

	greetUsers()

	for {

		firstName, secondName, email, tickets := getUserInput()

		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, secondName, email, tickets)

		if isValidName && isValidEmail && isValidTicket {

			bookTickets(firstName, secondName, tickets)

			firstNames := getFirstNames()

			fmt.Printf("Congratulations %v, your tickets are booked. There are %v tickets remaining.\n", firstName, remainingCoupons)

			fmt.Printf("These are the bookings: %v \n", firstNames)

			if remainingCoupons == 0 {
				fmt.Println("We have distributed all the coupons for today. Please come back tomorrow.")
				break
			}
		} else {
			fmt.Println("Incorrect input")
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v coupon booking.", foodCompany)
	fmt.Printf("\nWe have %v coupons left out of %v\n", remainingCoupons, totalCoupons)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var secondName string
	var email string
	var tickets uint
	fmt.Println("Please enter first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter second name:")
	fmt.Scan(&secondName)

	fmt.Println("Please enter your email")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets:")
	fmt.Scan(&tickets)

	return firstName, secondName, email, tickets
}

func validateUserInput(firstName string, secondName string, email string, tickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(secondName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := tickets > 0 && tickets <= remainingCoupons
	return isValidName, isValidEmail, isValidTicket
}

func bookTickets(firstName string, secondName string, tickets uint) {
	remainingCoupons = remainingCoupons - tickets
	bookings = append(bookings, firstName+" "+secondName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}
