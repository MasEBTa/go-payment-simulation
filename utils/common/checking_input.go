package common

import "regexp"

func isValidEmail(email string) bool {
	// Pola regex untuk memeriksa format email
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailPattern, email)
	return match
}