package vo

import "fmt"

//HDY123456
//
//HDY    : 3 digit static public holiday id identifier
//123456 : 6 digit random uppercase alphanumeric

const prefix = "HDY"

type PublicHolidayID string

func NewPublicHolidayID(random6Char string) (PublicHolidayID, error) {
	if len(random6Char) != 6 {
		return "", fmt.Errorf("must random 6 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random6Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	return PublicHolidayID(fmt.Sprintf("%s%s", prefix, random6Char)), nil
}

func (publicHolidayId PublicHolidayID) Validate() error {

	if len(publicHolidayId) != 9 {
		return fmt.Errorf("public holiday id must 9 characters")
	}

	if publicHolidayId[0:3] != prefix {
		return fmt.Errorf("public holiday id must started with %s %s", prefix, publicHolidayId[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(publicHolidayId[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", publicHolidayId[4:])
	}

	return nil
}
