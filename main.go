package main

import (
	"fmt"
	"strings"
)

func main() {
	const foodCompany string = "Shri Annapoorneshwari Bhojana"
	const totalCoupons uint = 200
	var remainingCoupons uint = 200
	var bookings []string

	fmt.Printf("Welcome to %v coupon booking.", foodCompany)
	fmt.Printf("\nWe have %v coupons left out of %v\n", remainingCoupons, totalCoupons)

	for {
		var firstName string
		var secondName string
		var tickets uint
		fmt.Println("Please enter first name:")
		fmt.Scan(&firstName)
		fmt.Println("Please enter second name:")
		fmt.Scan(&secondName)

		fmt.Println("Please enter number of tickets:")
		fmt.Scan(&tickets)

		if remainingCoupons < tickets {
			fmt.Printf("You cannot book %v coupons when only %v coupons are available. \n", tickets, remainingCoupons)
			break
		}

		remainingCoupons = remainingCoupons - tickets
		bookings = append(bookings, firstName+" "+secondName)

		firstNames := []string{}
		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstNames = append(firstNames, name[0])
		}

		fmt.Printf("Congratulations %v, your tickets are booked. There are %v tickets remaining.\n", firstName, remainingCoupons)

		fmt.Printf("These are the bookings: %v \n", firstNames)

		if remainingCoupons == 0 {
			fmt.Println("We have distributed all the coupons for today. Please come back tomorrow.")
			break
		}
	}

}
