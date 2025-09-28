package enum

type OrderCalculationDetailEnum string

const (
	StyleEnum = OrderCalculationDetailEnum("bold")

	LabelIMDEstimatedShippingCostEnumEN = OrderCalculationDetailEnum("Shipping Cost: (Instant) ")
	LabelIMDEstimatedShippingCostEnumID = OrderCalculationDetailEnum("Biaya Pengiriman: (Instan) ")

	LabelIMDDetailEstimatedShippingCostEnumEN = OrderCalculationDetailEnum("Shipping Cost Instant: (Destination %s) ")
	LabelIMDDetailEstimatedShippingCostEnumID = OrderCalculationDetailEnum("Biaya pengiriman Instant: (Tujuan %s) ")

	LabelEstimatedShippingCostEnumEN = OrderCalculationDetailEnum("Shipping Cost ")
	LabelEstimatedShippingCostEnumID = OrderCalculationDetailEnum("Biaya Kirim ")
	KeyEstimatedShippingCostEnum     = OrderCalculationDetailEnum("estimated_shipping_cost")

	LabelFoodSafeGuaranteeEnumEN = OrderCalculationDetailEnum("Food Safe Guarantee")
	LabelFoodSafeGuaranteeEnumID = OrderCalculationDetailEnum("Jaminan Makanan Aman")
	KeyFoodSafeGuaranteeEnum     = OrderCalculationDetailEnum("additional_cost_fsg")

	LabelSpecialHandlingFeeEnumEN = OrderCalculationDetailEnum("Special Handling Fee")
	LabelSpecialHandlingFeeEnumID = OrderCalculationDetailEnum("Biaya Penanganan Khusus")
	KeySpecialHandlingFeeEnum     = OrderCalculationDetailEnum("handling_fee")

	LabelSurchargeEnumEN = OrderCalculationDetailEnum("Surcharge")
	LabelSurchargeEnumID = OrderCalculationDetailEnum("Tambahan Biaya")
	KeySurchargeEnum     = OrderCalculationDetailEnum("surcharge")

	LabelInsuranceEnumEN = OrderCalculationDetailEnum("Insurance")
	LabelInsuranceEnumID = OrderCalculationDetailEnum("Asuransi")
	KeyInsuranceEnum     = OrderCalculationDetailEnum("insurance")

	LabelExtraPackagingEnumEN = OrderCalculationDetailEnum(" item | ")
	LabelExtraPackagingEnumID = OrderCalculationDetailEnum(" item | ")
	KeyExtraPackagingEnum     = OrderCalculationDetailEnum("extra_packaging_cost")

	LabelExtraBagEnumEN = OrderCalculationDetailEnum(" bag | ")
	LabelExtraBagEnumID = OrderCalculationDetailEnum(" tas | ")
	KeyExtraBagEnum     = OrderCalculationDetailEnum("extra_bag_cost")

	LabelDonationEnumEN = OrderCalculationDetailEnum(" item | ")
	LabelDonationEnumID = OrderCalculationDetailEnum(" item | ")
	KeyDonationEnum     = OrderCalculationDetailEnum("donation")

	LabelDiscountEnumEN = OrderCalculationDetailEnum("Discount")
	LabelDiscountEnumID = OrderCalculationDetailEnum("Diskon")
	KeyDiscountEnum     = OrderCalculationDetailEnum("discount_price")

	LabelAdminFeeEnumEN = OrderCalculationDetailEnum("Admin fee")
	LabelAdminFeeEnumID = OrderCalculationDetailEnum("Biaya Admin")
	KeyAdminFeeEnum     = OrderCalculationDetailEnum("admin_fee")

	LabelCOPFeeEnumEN = OrderCalculationDetailEnum("COP fee")
	LabelCOPFeeEnumID = OrderCalculationDetailEnum("Biaya COP")
	KeyCOPFeeEnum     = OrderCalculationDetailEnum("cop_fee")

	LabelTotalPaymentEnumEN = OrderCalculationDetailEnum("Total Payment")
	LabelTotalPaymentEnumID = OrderCalculationDetailEnum("Biaya Total")
	KeyTotalPaymentEnum     = OrderCalculationDetailEnum("total_payment")
)
