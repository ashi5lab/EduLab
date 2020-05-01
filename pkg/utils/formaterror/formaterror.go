package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors.New("Nickname Already Taken")
	}

	if strings.Contains(err, "record not found") {
		return errors.New("Email Not registered")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Details")
	}
	return errors.New("Incorrect Details")
}
