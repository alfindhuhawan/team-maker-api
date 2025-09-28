package vo

import "fmt"

//CRT0123AB
//
//CRT    : 3 digit static Cart id identifier
//0123AB : 6 digit random uppercase alphanumeric
//
//can hold max 36^6 Cart per system
//by default everyone is a Cart

type CartID string

func NewCartID(buyerID BuyerID) (CartID, error) {
	random6Char := buyerID[3:]
	return CartID(fmt.Sprintf("CRT%s", random6Char)), nil
}

func (b CartID) Validate() error {

	if len(b) != 9 {
		return fmt.Errorf("cart id must 9 characters")
	}

	if b[0:3] != "CRT" {
		return fmt.Errorf("cart id must started with CRT %s", b[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(b[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", b[4:])
	}

	return nil
}
