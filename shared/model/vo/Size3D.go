package vo

import (
	"fmt"
)

type DisplayEmptySize3D [3]string
type Size3D [3]int

func NewSize3D(length, width, height int) Size3D {
	return [3]int{length, width, height}
}

func NewDisplayEmptySize3D(length, width, height string) DisplayEmptySize3D {
	return [3]string{length, width, height}
}

func (s Size3D) Validate() error {

	// TODO must not 0 or negative

	var invalid bool = false

	if len(s) != 3 {
		invalid = true
	}

	for _, v := range s {
		if v < 0 { // CANT INPUT MINUS, NOW CHANGE TO 0
			invalid = true
			break
		}
	}

	if invalid {
		return fmt.Errorf("invalid size format errrr")
	}

	return nil
}

func (s Size3D) String() string {
	return fmt.Sprintf("%dx%dx%d", s[0], s[1], s[2])
}
