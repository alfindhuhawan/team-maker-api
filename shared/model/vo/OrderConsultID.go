package vo

import (
	"fmt"
	"time"
)

//PXL    : 3 digit static order id identifier
//220119 : 2 digit year, 2 digit month, 2 digit date
//150807 : 2 digit hour, 2 digit minute, 2 digit second
//X | G  : identifier for login user (X) and guest user (G)
//0123AB : RandomString

type OrderConsultID string

func NewOrderConsultID(now time.Time, paidBy string, randomString string) (OrderConsultID, error) {
	labelType := "X"
	if paidBy == "" {
		labelType = "G"
	}

	return OrderConsultID(fmt.Sprintf("PXM%s%s%s", now.Format("060102150405"), labelType, randomString)), nil
}

func (n OrderConsultID) Validate() error {

	if len(n) != 22 {
		return fmt.Errorf("order id must 22 characters")
	}

	if n[0:3] != "PXM" && n[0:3] != "C2B" {
		return fmt.Errorf("buyer id must started with PXM / C2B not %s", n[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(n[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", n[4:])
	}

	return nil
}

// func (n OrderID) ExtractBuyerID() BuyerID {
// 	return BuyerID(fmt.Sprintf("BYR%s", n[16:22]))
// }

func (n OrderConsultID) String() string {
	return string(n)
}
