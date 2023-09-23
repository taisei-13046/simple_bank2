package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("length must be between %d and %d", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(username string) error {
	if err := ValidateString(username, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(username) {
		return fmt.Errorf("username must be alphanumeric")
	}
	return nil
}

func ValidatePassword(password string) error {
	if err := ValidateString(password, 6, 100); err != nil {
		return err
	}
	return nil
}

func ValidateEmail(email string) error {
	if err := ValidateString(email, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

func ValidateUserFullName(fullName string) error {
	if err := ValidateString(fullName, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(fullName) {
		return fmt.Errorf("username must be alphanumeric")
	}
	return nil
}
