package vo

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
)

type PhoneNumber string

func (n PhoneNumber) Validate(field string) error {
	if len(n) < 10 {
		return fmt.Errorf("Nomor %s minimal 10 karakter", field)
	}

	if len(n) > 14 {
		return fmt.Errorf("Nomor %s maksimal 14 karakter", field)
	}

	if n[0:3] != "628" && n[0:2] != "08" {
		return fmt.Errorf("Nomor %s tidak sesuai format", field)
	}

	if n[0:3] == "628" && len(n) < 11 {
		return fmt.Errorf("Nomor %s minimal 11 karakter", field)
	}

	_, err := phonenumbers.Parse(n.String(), "ID")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Nomor %s tidak valid", field)
	}

	return nil
}

func (n PhoneNumber) String() string {

	// TODO convert to some format

	return string(n)
}

func FormatIDN(phoneNumber string) string {
	new := ""

	if phoneNumber != "" {
		// CONVERT 0 TO 62
		if phoneNumber[0:1] == "0" {
			new = "62" + string(phoneNumber[1:])
		} else {
			new = string(phoneNumber)
		}
	}

	return string(new)
}

func ChangeToZero(phoneNumber string) string {
	new := ""

	// CONVERT 0 TO 62
	if phoneNumber[0:2] == "62" {
		new = "0" + string(phoneNumber[2:])
	} else {
		new = string(phoneNumber)
	}

	return string(new)
}
