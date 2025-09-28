package vo

import (
	"fmt"
)

// ADM123456
//
//ADM    : 3 digit static admin id identifier
//0123AB : 6 digit random uppercase alphanumeric
//
//can hold max 36^6 admin per system

type AdminID string

func NewAdminID(random6Char string) (AdminID, error) {
	if len(random6Char) != 6 {
		return "", fmt.Errorf("must random 6 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random6Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	return AdminID(fmt.Sprintf("ADM%s", random6Char)), nil
}

func (b AdminID) Validate() error {

	if len(b) != 9 {
		return fmt.Errorf("admin id must 9 characters")
	}

	if b[0:3] != "ADM" {
		return fmt.Errorf("admin id must started with ADM %s", b[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(b[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", b[4:])
	}

	return nil
}
