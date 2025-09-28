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

type SellerID string

func NewSellerID(buyerID BuyerID) (SellerID, error) {
	random6Char := buyerID[3:]
	return SellerID(fmt.Sprintf("SLR%s", random6Char)), nil
}

func (n SellerID) Validate() error {

	if len(n) != 9 {
		return fmt.Errorf("seller id must 9 characters")
	}

	if n[0:3] != "SLR" {
		return fmt.Errorf("seller id must started with SLR %s", n[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(n[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", n[4:])
	}

	return nil
}

func (n SellerID) String() string {
	return string(n)
}
