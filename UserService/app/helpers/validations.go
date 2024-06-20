package helpers

func ValidatePassword(password string) bool {
	return len([]rune(password)) >= 4
}

func ValidateEmail(email string) bool {
	return true
}
