package enum

type ErrorFromEnum string

const (
	ValidationErrorEnum                = ErrorFromEnum("Validation")
	CheckCoverageErrorEnum             = ErrorFromEnum("Check Coverage")
	GroupDaysValidationErrorEnum       = ErrorFromEnum("Group Days Validation")
	ShipmentOrderLimitErrorEnum        = ErrorFromEnum("Shipment Order Limit")
	GetMemberByPhoneErrorEnum          = ErrorFromEnum("Get Member By Phone")
	CalculateShippingCostErrorEnum     = ErrorFromEnum("Calculate Shipping Cost")
	GetAdminFeeErrorEnum               = ErrorFromEnum("Get Admin Fee")
	GetPricePackageErrorEnum           = ErrorFromEnum("Get Price Package")
	GetPackageAddonBySizeMainErrorEnum = ErrorFromEnum("Get Package Addon By Size")
	CalculteInsuranceFeeMainErrorEnum  = ErrorFromEnum("Calculate Insurance Fee")
	PromoValidationPGWMainErrorEnum    = ErrorFromEnum("Promo Validation PGW")
	PromoAutoApplyPGWMainErrorEnum     = ErrorFromEnum("Promo Auto Apply Voucher PGW")
	CalculateDonationMainErrorEnum     = ErrorFromEnum("Calculate Donation")
	PaxelRecycleValidationErrorEnum    = ErrorFromEnum("Paxel Recycle Validation")
	GetPriceTimeMultidropErrorEnum     = ErrorFromEnum("Get Price Time Instant Multidrop")
	CalculteOtherShipmentFeeErrorEnum  = ErrorFromEnum("Calculate Other Shipment Fee")
	CalculteCOPFeeErrorEnum            = ErrorFromEnum("Calculate COP Fee")
	GetPriceTimeErrorEnum              = ErrorFromEnum("Get Price Time")
)
