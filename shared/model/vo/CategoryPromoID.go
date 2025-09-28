package vo

import (
	"fmt"
	"time"
)

type CategoryPromoID string

func NewCategoryPromoID(random6Char string) (CategoryPromoID, error) {
	return CategoryPromoID(fmt.Sprintf("CPR%s", random6Char)), nil
}

func ParseStringToDateTime(dateString string) (time.Time, error) {
	result, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return time.Time{}, err
	}
	return result, nil
}

func (id CategoryPromoID) String() string {
	return string(id)
}
