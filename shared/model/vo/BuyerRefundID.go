package vo

import (
	"fmt"
	"time"
)

// RFN220418112233X123456
//
//RFN    : 3 digit static buyer refund id identifier
//220418 : 6 digit Y-m-d
//112233 : 6 digit H:i:s
//X                      : separator
//123456 : BuyerID

type BuyerRefundID string

func NewBuyerRefundID(now time.Time, buyerID BuyerID) (BuyerRefundID, error) {
	return BuyerRefundID(fmt.Sprintf("RFN%sX%s", now.Format("060102150405"), buyerID[3:])), nil
}

func (b BuyerRefundID) Validate() error {

	if len(b) != 22 {
		return fmt.Errorf("buyer refund id must 22 characters")
	}

	if b[0:3] != "RFN" {
		return fmt.Errorf("buyer refund id must started with RFN %s", b[0:3])
	}

	if !regexUppercaseAlphaNumeric.MatchString(string(b[4:])) {
		return fmt.Errorf("last 6 digit must uppercase alphanumeric %s", b[4:])
	}

	return nil
}
