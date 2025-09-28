package vo

import (
	"fmt"
)

// ART123456
//
//ART    : 3 digit static article id identifier
//0123AB : 6 digit random uppercase alphanumeric
//
//can hold max 36^6 article per system

type ArticleID string

func NewArticleID(random6Char string) (ArticleID, error) {
	if len(random6Char) != 6 {
		return "", fmt.Errorf("must random 6 digit char")
	}

	if !regexUppercaseAlphaNumeric.MatchString(random6Char) {
		return "", fmt.Errorf("must uppercase alphanumeric")
	}

	return ArticleID(fmt.Sprintf("ART%s", random6Char)), nil
}

func (b ArticleID) Validate() error {

	if len(b) != 9 {
		return fmt.Errorf("article id must 9 characters")
	}

	if b[0:3] != "ART" {
		return fmt.Errorf("article id must started with ART %s", b[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(b[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", b[4:])
	}

	return nil
}
