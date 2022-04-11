package helper

import "strings"

func ValidateUserInput(remainingCoupons uint, firstName string, secondName string, email string, tickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(secondName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := tickets > 0 && tickets <= remainingCoupons
	return isValidName, isValidEmail, isValidTicket
}
