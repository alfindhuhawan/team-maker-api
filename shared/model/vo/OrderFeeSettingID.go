package vo

import (
	"fmt"
)

// OFS123456
//
//FES    : 3 digit static admin id identifier
//0123AB : 6 digit random uppercase alphanumeric
//
//can hold max 36^6 admin per system

type OrderFeeSettingID string

func NewOrderFeeSettingID(random6Char string) (OrderFeeSettingID, error) {
	if len(random6Char) != 6 {
		return "", fmt.Errorf("must random 6 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random6Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	//return OrderFeeSettingID(fmt.Sprintf("FES%s", random6Char)), nil //if separated from global setting
	return OrderFeeSettingID(fmt.Sprintf(random6Char)), nil
}

func (b OrderFeeSettingID) Validate() error {

	if len(b) != 9 {
		return fmt.Errorf("admin id must 9 characters")
	}

	if b[0:3] != "FES" {
		return fmt.Errorf("admin id must started with FES %s", b[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(b[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", b[4:])
	}

	return nil
}
