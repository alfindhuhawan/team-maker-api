package vo

import "fmt"

type Percentage float64

func (p Percentage) Validate() error {

	if p < 0 || p > 100 {
		return fmt.Errorf("percentage x value must 0 < x <= 100")
	}

	return nil

}
