package vo

import "fmt"

type Rating float32

func (r Rating) Validate() error {
	if r <= 0 || r > 5 {
		return fmt.Errorf("rating must 1-5")
	}
	return nil
}

func NewRating(x float32) (Rating, error) {
	if x <= 0 || x > 5 {
		return 0, fmt.Errorf("rating must 1-5")
	}
	return 0, nil
}
