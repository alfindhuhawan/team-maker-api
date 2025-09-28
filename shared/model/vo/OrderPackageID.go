package vo

import (
	"fmt"
)

type OrderPackageID string

func NewOrderPackageID(orderSellerID OrderSellerID, index int) (OrderPackageID, error) {
	return OrderPackageID(fmt.Sprintf("%s-%d", orderSellerID, index)), nil
}

func (o OrderPackageID) String() string {
	return string(o)
}
