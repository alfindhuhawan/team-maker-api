package vo

import (
	"fmt"
	"time"
)

// INC220119150807G0123AB

//INC    : 3 digit static id identifier
//220119 : 2 digit year, 2 digit month, 2 digit date
//150807 : 2 digit hour, 2 digit minute, 2 digit second
//X | G  : identifier for login user (X) and guest user (G)
//0123AB : 6 digit random uppercase alphanumeric

type SellerIncomeID string

func NewSellerIncomeID(now time.Time, random6Char string) (SellerIncomeID, error) {
	if len(random6Char) != 6 {
		return "", fmt.Errorf("must random 6 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random6Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	return SellerIncomeID(fmt.Sprintf("INC%sX%s", now.Format("060102150405"), random6Char)), nil
}

func (n SellerIncomeID) Validate() error {

	if len(n) != 22 {
		return fmt.Errorf("id must 22 characters")
	}

	if n[0:3] != "INC" {
		return fmt.Errorf("id must started with INC %s", n[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(n[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", n[4:])
	}

	return nil
}
