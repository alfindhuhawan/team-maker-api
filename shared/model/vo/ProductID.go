package vo

import (
	"fmt"
)

// WARNING READ IT BEFORE CHANGE THE FORMAT !! IF YOU CHANGE THE FORMAT SO YOU MUST CHANGE THE LOGIC TO GET LAST COUNTER ON IMPL PRODUCT FUNC "FindLastProductIDCounter"

//PRD0123AB008ABCDEF
//
//PRD    : 3 digit static product id identifier
//0123AB : 6 digit seller id (without SLR prefix)
//008ABCDEF    : 9 random char
//

type ProductID string

func NewProductID(sellerID SellerID, random9Char string) (ProductID, error) {
	return ProductID(fmt.Sprintf("PRD%s%s", sellerID[3:], random9Char)), nil
}

func (p ProductID) Validate() error {

	if len(p) != 12 && len(p) != 18 {
		return fmt.Errorf("product id must 12 characters or 18 characters")
	}

	if p[0:3] != "PRD" {
		return fmt.Errorf("product id must started with PRD %s", p[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(p[4:9])) {
		return fmt.Errorf("next 6 digit must uppercase alphanumeric %s", p[4:9])
	}

	// if !regexNumeric.MatchString(string(p[10:12])) {
	// 	return fmt.Errorf("last 3 digit must numeric %s", p[10:12])
	// }

	return nil
}

func (p ProductID) ExtractSellerID() SellerID {
	return SellerID(fmt.Sprintf("SLR%s", p[3:9]))
}

func (p ProductID) String() string {
	return string(p)
}
