package vo

import "fmt"

//PRD0123ABCATMWI42G
//PRD    : 3 digit static product id identifier
//0123AB : 6 digit product id (without PRD prefix)
//CAT    : 3 digit static category id identifier
//MWI42G : 6 digit category id (without CAT prefix)

type ProductCategoryID string

func NewProductCategoryID(productID ProductID, categoryID CategoryID) (ProductCategoryID, error) {
	return ProductCategoryID(fmt.Sprintf(string(productID) + string(categoryID))), nil
}
