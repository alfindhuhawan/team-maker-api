package vo

import (
	"fmt"
)

//PXM220119150807X0123AB-0123AB
//
//PXM220119150807X0123AB : Parent OrderID
//-                      : separator
//0123AB : Seller OrderID without (SLR prefix)

type OrderSellerID string

func NewOrderSellerID(orderID OrderID, sellerID SellerID) (OrderSellerID, error) {
	return OrderSellerID(fmt.Sprintf("%s-%s", orderID, sellerID[3:])), nil
}

func (n OrderSellerID) Validate() error {

	// TODO validate

	return nil
}

func (n OrderSellerID) String() string {
	return string(n)
}
