package vo

import "fmt"

//CAT0123AB
//
//CAT    : 3 digit static category id identifier
//0123AB : 6 digit random uppercase alphanumeric
//
//can hold max 36^6 category per system
//by default everyone is a category

type CategoryID string

func NewCategoryID(random3Char string) (CategoryID, error) {
	if len(random3Char) != 3 {
		return "", fmt.Errorf("must random 3 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random3Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	return CategoryID(fmt.Sprintf("CAT%s", random3Char)), nil
}

func NewSubCategoryID(parentID CategoryID, random4Char string) (CategoryID, error) {
	if len(random4Char) != 4 {
		return "", fmt.Errorf("must random 4 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random4Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	return CategoryID(fmt.Sprintf("%s-%s", parentID, random4Char)), nil
}

func (b CategoryID) Validate() error {

	if len(b) != 9 {
		return fmt.Errorf("category id must 9 characters")
	}

	if b[0:3] != "CAT" {
		return fmt.Errorf("category id must started with CAT %s", b[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(b[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", b[4:])
	}

	return nil
}
