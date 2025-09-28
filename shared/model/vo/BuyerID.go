package vo

import (
	"fmt"
)

//BYR0123AB
//
//BYR    : 3 digit static buyer id identifier
//0123AB : 6 digit random uppercase alphanumeric
//
//can hold max 36^6 buyer per system
//by default everyone is a Buyer

type BuyerID string

func NewBuyerID(random6Char string) (BuyerID, error) {
	if len(random6Char) != 6 {
		return "", fmt.Errorf("must random 6 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random6Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	return BuyerID(fmt.Sprintf("BYR%s", random6Char)), nil
}

func (b BuyerID) Validate() error {

	if len(b) != 9 {
		return fmt.Errorf("buyer id must 9 characters")
	}

	if b[0:3] != "BYR" {
		return fmt.Errorf("buyer id must started with BYR %s", b[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(b[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", b[4:])
	}

	return nil
}
