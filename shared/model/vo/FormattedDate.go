package vo

import "time"

type FormattedDate string

func (f FormattedDate) Validate() error {
	_, err := time.Parse("060102", string(f))
	return err
}
