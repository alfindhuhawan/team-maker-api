package vo

import (
	"fmt"
)

//SLR0123AB
//
//SLR    : 2 digit static seller id identifier
//0123AB : 6 digit buyer id (without BYR prefix)
//
//The number of seller is less or equal than the buyer number
//Seller basically is a Buyer who activate the Seller Mode

type BannerID string

func NewBannerID(random6Char string) (BannerID, error) {
	if len(random6Char) != 6 {
		return "", fmt.Errorf("must random 6 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random6Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	return BannerID(fmt.Sprintf("BNR%s", random6Char)), nil
}
