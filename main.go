package main

import "fmt"

func main() {
	const foodCompany string = "Shri Annapoorneshwari Bhojana"
	const totalCoupons uint = 200
	var remainingCoupons uint = 200

	fmt.Printf("Welcome to %v coupon booking.", foodCompany)
	fmt.Printf("\nWe have %v coupons left out of %v\n", remainingCoupons, totalCoupons)

	var name string
	var tickets uint
	fmt.Println("Please enter name:")
	fmt.Scan(&name)

	fmt.Println("Please enter number of tickets:")
	fmt.Scan(&tickets)

	remainingCoupons = remainingCoupons - tickets

	fmt.Printf("\nCongratulations %v, your tickets are booked. There are %v tickets remaining.\n", name, remainingCoupons)
}
