package vo

import (
	"fmt"
)

//REV220119150807X0123AB-0123AB
//REV    : 3 digit static admin id identifier
//220119150807X0123AB : Parent OrderID
//-                      : separator
//0123AB : Seller OrderID without (SLR prefix)

type ReviewID string

func NewReviewID(orderSellerID OrderSellerID, sellerID SellerID, productID ProductID) (ReviewID, error) {
	orderID := orderSellerID[3:]

	return ReviewID(fmt.Sprintf("REV%s-%s", orderID, productID)), nil
}

func (n ReviewID) Validate() error {

	// TODO validate

	return nil
}
