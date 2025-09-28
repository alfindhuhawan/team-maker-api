package util

import (
	"fmt"
	"team-maker-api/shared/model/vo"
)

func GetDiscount(priceRegular vo.Money, priceSale vo.Money) (vo.Percentage, error) {
	var discount vo.Percentage

	if priceRegular < priceSale {
		msg := fmt.Sprintf("harga diskon %s tidak boleh lebih besar dari %s", priceSale, priceRegular)
		return 0, fmt.Errorf(msg)
	}

	if priceRegular == 0 && priceSale == 0 {
		return 0, nil
	}

	discount = ((vo.Percentage(priceRegular) - vo.Percentage(priceSale)) * 100) / vo.Percentage(priceRegular)

	return discount, nil
}
