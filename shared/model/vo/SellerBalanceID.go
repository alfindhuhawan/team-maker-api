package vo

import (
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid"
)

// BLC220418112233X123456
//
//BLC    : 3 digit static Seller balance id identifier
//220418 : 6 digit Y-m-d
//112233 : 6 digit H:i:s
//X      : separator
//123456 : SellerID

type SellerBalanceID string

const prefixSellerBalanceID = "BLC"

func NewSellerBalanceID(now time.Time, sellerID SellerID) (SellerBalanceID, error) {

	err := SellerID.Validate(sellerID)
	if err != nil {
		return "", err
	}

	// randomString := util.GenerateID(3)
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	randomString, err := gonanoid.Generate(alphabet, 3)
	if err != nil {
		return "", fmt.Errorf("error create random string")
	}

	// CHANGE TIME TO JAKARTA
	location, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(location)

	return SellerBalanceID(fmt.Sprintf("%s%s%sX%s", prefixSellerBalanceID, timeNow.Format("060102150405"), randomString, sellerID[3:])), nil
}

func (n SellerBalanceID) Validate() error {

	if len(n) >= 22 && len(n) <= 25 {
		return fmt.Errorf("id must 22 or 25 characters")
	}

	if n[0:3] != prefixSellerBalanceID {
		return fmt.Errorf("id must started with %s %s", prefixSellerBalanceID, n[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(n[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", n[4:])
	}

	return nil
}
