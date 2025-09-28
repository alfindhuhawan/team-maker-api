package vo

import (
	"fmt"
)

type PaymentHistoryID string

func NewPaymentHistoryID(orderID OrderID, code string) (PaymentHistoryID, error) {

	err := orderID.Validate()
	if err != nil {
		return "", err
	}

	return PaymentHistoryID(fmt.Sprintf("%s%s", orderID, code)), nil
}
