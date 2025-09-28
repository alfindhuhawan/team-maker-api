package vo

import "fmt"

//BYR0123ABPRD0123AB008
//
//BYR    : 3 digit static buyer id identifier
//0123AB : 6 digit random uppercase alphanumeric
//
//PRD    : 3 digit static product id identifier
//0123AB : 6 digit seller id (without SLR prefix)
//008    : 3 digit sequensial number
//
//can hold max : 999 productfavorite

type ProductFavoriteID string

func NewProductFavoriteID(buyerID BuyerID, productID ProductID) (ProductFavoriteID, error) {
	return ProductFavoriteID(fmt.Sprintf(string(buyerID) + string(productID))), nil
}

// func (p ProductID) Validate() error {

// 	if len(p) != 12 {
// 		return fmt.Errorf("product id must 12 characters")
// 	}

// 	if p[0:3] != "PRD" {
// 		return fmt.Errorf("product id must started with PRD %s", p[0:3])
// 	}

// 	if !regexUppercaseAlphaNumeric.MatchString(string(p[4:9])) {
// 		return fmt.Errorf("next 6 digit must uppercase alphanumeric %s", p[4:9])
// 	}

// 	if !regexNumeric.MatchString(string(p[10:12])) {
// 		return fmt.Errorf("last 3 digit must numeric %s", p[10:12])
// 	}

// 	return nil
// }

// func (p ProductID) ExtractSellerID() SellerID {
// 	return SellerID(fmt.Sprintf("SLR%s", p[3:9]))
// }
