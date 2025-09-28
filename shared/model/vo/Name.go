package vo

import (
	"fmt"
)

type Name string

func (n Name) Validate() error {

	// TODO not allowed emoji or any other validation

	if n == "" {
		return fmt.Errorf("name must not empty")
	}

	return nil
}

func (n Name) String() string {
	return string(n)
}
