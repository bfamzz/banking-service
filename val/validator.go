package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
)

func ValidateStringLength(data string, minimum, maximum int) error {
	if len(data) < minimum || len(data) > maximum {
		return fmt.Errorf("string length should be between %d - %d", minimum, maximum)
	}
	return nil
}

func ValidateUsername(username string) error {
	if err := ValidateStringLength(username, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(username) {
		return fmt.Errorf("username must only contain lowercase letters, digits and/or underscore")
	}
	return nil
}

func ValidatePassword(password string) error {
	return ValidateStringLength(password, 6, 100)
}

func ValidateEmail(email string) error {
	if err := ValidateStringLength(email, 6, 200); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("%s is not a valid email address", email)
	}
	return nil
}

func ValidateFullName(fullName string) error {
	if err := ValidateStringLength(fullName, 3, 100); err != nil {
		return err
	}

	if !isValidFullName(fullName) {
		return fmt.Errorf("fullname must only contain letters or spaces")
	}
	return nil
}