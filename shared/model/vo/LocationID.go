package vo

import (
	"fmt"
)

// ADR123456001
//
//ADR    : 3 digit static Location id identifier
//123456 : 6 digit Buyer ID
//001 : 3 digit random uppercase alphanumeric

type LocationID string

func NewLocationID(random3Char string, referenceID string) (LocationID, error) {
	if len(random3Char) != 3 {
		return "", fmt.Errorf("must random 3 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random3Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	refStringID := referenceID[4:]

	return LocationID(fmt.Sprintf("ADR%s%s", refStringID, random3Char)), nil
}

func (b LocationID) Validate() error {

	if len(b) != 12 {
		return fmt.Errorf("Address id must 12 characters")
	}

	if b[0:3] != "ADR" {
		return fmt.Errorf("Address id must started with ADR %s", b[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(b[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", b[4:])
	}

	return nil
}
