package helpers

import "errors"

func ValidatePassword(password string) error {
	if len([]rune(password)) >= 4 {
		return nil
	}
	return errors.New("your password is too short")
}

func ValidateEmail(email string) error {

	return nil
}
