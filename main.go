package main

import (
	"fmt"
	"golang-learning/helper"
	"time"
)

const foodCompany string = "Shri Annapoorneshwari Bhojana"
const totalCoupons uint = 200

var remainingCoupons uint = 200
var bookings = make([]UserData, 0)

type UserData struct {
	firstName  string
	secondName string
	email      string
	tickets    uint
}

func main() {

	greetUsers()

	for {

		firstName, secondName, email, tickets := getUserInput()

		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(remainingCoupons, firstName, secondName, email, tickets)

		if isValidName && isValidEmail && isValidTicket {

			bookTickets(firstName, secondName, email, tickets)
			go sendTickets(firstName, secondName, email, tickets)

			firstNames := getFirstNames()

			fmt.Printf("Congratulations %v, your tickets are booked. There are %v tickets remaining.\n", firstName, remainingCoupons)

			fmt.Printf("These are the bookings: %v \n", firstNames)
			fmt.Printf("These are the complete bookings: %v \n", bookings)

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

func bookTickets(firstName string, secondName string, email string, tickets uint) {
	remainingCoupons = remainingCoupons - tickets
	var userData = UserData{
		firstName:  firstName,
		secondName: secondName,
		email:      email,
		tickets:    tickets,
	}
	bookings = append(bookings, userData)
}

func sendTickets(firstName string, secondName string, email string, tickets uint) {
	time.Sleep(10 * time.Second)
	var msg = fmt.Sprintf("%v tickets sending to %v %v", tickets, firstName, secondName)
	fmt.Printf("###################################################\n")
	fmt.Printf("Sending tickets... \n %v \n to email %v\n", msg, email)
	fmt.Printf("###################################################\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
