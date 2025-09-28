package vo

import (
	"fmt"
)

//RFD123ABC

type OrderSellerRefundID string

func NewOrderSellerRefundID(Sequence string) (OrderSellerRefundID, error) {
	return OrderSellerRefundID(fmt.Sprintf("RFD%s", Sequence)), nil
}

func (n OrderSellerRefundID) Validate() error {

	// TODO validate

	return nil
}

func (n OrderSellerRefundID) String() string {
	return string(n)
}
