package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remaningTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") //strings.Contains ile email içinde @ karakteri var mı diye bakar.
	isValidTicketNumber := userTickets > 0 && userTickets <= remaningTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
