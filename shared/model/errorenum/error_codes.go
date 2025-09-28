package errorenum

// GLOBAL ERROR
const (
	parsingFailedEN     = "Parsing %s failed >> %s"
	parsingFailedID     = "Gagal Melakukan Parsing %s >> %s"
	outOfServiceTime    = "OUT_OF_SERVICE_TIME %s"
	internalError       = "Internal error >> %s"
	failureConnectionEN = "failure connection"
	failureConnectionID = "koneksi gagal"
)

const (
	// ERROR SYSTEM
	SomethingError    ErrorType = "ER0000 something error"
	FailureConnection ErrorType = "ER5001 terjadi kesalahan koneksi"
	ParsingData       ErrorType = "ER5002 terjadi kesalahan data"
	ResponseNotOk     ErrorType = "ER5003 terjadi gangguan pada sistem internal"
	RabbitMQError     ErrorType = "ERRBTMQ0001 Terjadi kesalahan pada koneksi, silakan coba beberapa saat lagi."

	// ERROR GET PRICE TIME
	PriceTimeFailureConnectionEN ErrorType = "ERPT0000 " + failureConnectionEN + " (price time)"
	PriceTimeFailureConnectionID ErrorType = "ERPT0000 " + failureConnectionID + " (price time)"
	PriceTimeInternalError       ErrorType = "ERPT0001 " + internalError
	PriceTimeError               ErrorType = "ERPT0002 %s"

	// ERROR CHECK COVERAGE
	CheckCoverageFailureConnectionEN  ErrorType = "ERCC0000 " + failureConnectionEN + " (check coverage)"
	CheckCoverageFailureConnectionID  ErrorType = "ERCC0000 " + failureConnectionID + " (check coverage)"
	CheckCoverageInternalError        ErrorType = "ERCC0001 " + internalError
	CheckCoverageError                ErrorType = "ERCC0002 %s"
	CheckCoverageAddressCantBeEmptyEN ErrorType = "ERCC0001 Addresses must to be filled"
	CheckCoverageAddressCantBeEmptyID ErrorType = "ERCC0001 Alamat harus diisi"

	// ERROR VALIDATION
	ValidationCantBeEmptyEN            ErrorType = "ERV0001 %s cant be empty"
	ValidationCantBeEmptyID            ErrorType = "ERV0001 %s tidak boleh kosong"
	WrongFormatEmailEN                 ErrorType = "ERV0002 %s wrong format email"
	WrongFormatEmailID                 ErrorType = "ERV0002 %s format email salah"
	ValidationWrongFormatChannelSource ErrorType = "ERV0003 channel source not identify"

	// ERROR GROUP VALIDATION
	GVParsingFailedEN      ErrorType = "ERGV0001 " + parsingFailedEN
	GVParsingFailedID      ErrorType = "ERGV0001 " + parsingFailedID
	GVOutOfServiceTime     ErrorType = "ERGV0002 " + outOfServiceTime
	GVDTNotFoundEN         ErrorType = "ERGV0003 Delivery Time in Time Table not found"
	GVDTNotFoundID         ErrorType = "ERGV0003 Waktu Pengiriman di Time Table tidak di temukan"
	GVPriceTimeError       ErrorType = "ERGV0004 %s"
	GVShipmentSettingError ErrorType = "ERGV0005 %s"

	// ERROR SHIPMENT ORDER LIMIT
	SOLFailureConnectionEN ErrorType = "ERSOL0000 " + failureConnectionEN + " (shipment order limit)"
	SOLFailureConnectionID ErrorType = "ERSOL0000 " + failureConnectionID + " (shipment order limit)"
	SOLInternalError       ErrorType = "ERSOL0001 " + internalError
	SOLError               ErrorType = "ERSOL0003 %s"

	// ERROR SHIPMENT SETTINGS
	ShipmentSettingFailureConnectionEN ErrorType = "ERSS0000 " + failureConnectionEN + " (shipment setting)"
	ShipmentSettingFailureConnectionID ErrorType = "ERSS0000 " + failureConnectionID + " (shipment setting)"
	ShipmentSettingInternalError       ErrorType = "ERSS0001 " + internalError
	ShipmentSettingError               ErrorType = "ERSS0002 %s"

	// PHONE NUMBER VALIDATION
	PhoneNumberValidationFailureConnectionEN ErrorType = "ERPNV0000 " + failureConnectionEN + " (phone number validation)"
	PhoneNumberValidationFailureConnectionID ErrorType = "ERPNV0000 " + failureConnectionID + " (phone number validation)"
	PhoneNumberValidationInternalError       ErrorType = "ERPNV0001 " + internalError
	PhoneNumberValidationError               ErrorType = "ERPNV0002 %s"

	// GET PREFIX CLUSTER AREA CODE FOR PAXEL COUNTER
	GPCACFPCFailureConnectionEN ErrorType = "ERGPCACFPC0000 " + failureConnectionEN + " (get prefix cluster area code for paxel counter)"
	GPCACFPCFailureConnectionID ErrorType = "ERGPCACFPC0000 " + failureConnectionID + " (get prefix cluster area code for paxel counter)"
	GPCACFPCInternalError       ErrorType = "ERGPCACFPC0001 " + internalError
	GPCACFPCError               ErrorType = "ERGPCACFPC0002 %s"

	// GET NEAREST PAXEL COUNTERS MAIN
	GNPCMFailureConnectionEN ErrorType = "ERGNPCM0000 " + failureConnectionEN + " (get nearest paxel counters main)"
	GNPCMFailureConnectionID ErrorType = "ERGNPCM0000 " + failureConnectionID + " (get nearest paxel counters main)"
	GNPCMInternalError       ErrorType = "ERGNPCM0001 " + internalError
	GNPCMError               ErrorType = "ERGNPCM0002 %s"

	// ERROR DONATION CALCULATION
	DonationExternalConnectionEN ErrorType = "ERPT0000 " + failureConnectionEN + " (donation calculation)"
	DonationExternalConnectionID ErrorType = "ERPT0000 " + failureConnectionID + " (donation calculation)"
	DonationInternalError        ErrorType = "ERPT0001 " + internalError
	DonationError                ErrorType = "ERPT0002 %s"
	DonationIDMandatory          ErrorType = "ERGV0003 ID is mandatory"
	DonationFSCError             ErrorType = "ERGV0004 Fix Shiping Cost should be greater than 0"
	DonationNotFoundError        ErrorType = "ERGV0005 Donation not found"

	// ERROR OTHER FEE CALCULATION
	OFShipmentSettingError ErrorType = "EROF0000 %s"
	OFConvertStringError   ErrorType = "EROF0001 %s"

	// ERROR INSURANCE CALCULATION AND CHECK
	INShipmentSettingError ErrorType = "ERIN0000 %s"
	INConvertStringError   ErrorType = "ERIN0001 %s"
	INDefaultError         ErrorType = "ERIN0002 %s"
	INConvertIntError      ErrorType = "ERIN0003 %s"

	// ERROR ADMIN FEE CALCULATION
	AdminFeeConnectionEN    ErrorType = "ERAF0000 " + failureConnectionEN + " (get admin fee)"
	AdminFeeConnectionID    ErrorType = "ERAF0000 " + failureConnectionID + " (get admin fee)"
	AdminFeeParsingFailedEN ErrorType = "ERAF0001 " + parsingFailedEN + " (get admin fee)"
	AdminFeeParsingFailedID ErrorType = "ERAF0001 " + parsingFailedID + " (get admin fee)"
	AdminFeeResponseNotOK   ErrorType = "ERAF0002 " + ResponseNotOk + " %s"
	AdminFeeInternalError   ErrorType = "ERAF0003 %s"

	// ERROR ADMIN FEE CALCULATION
	ExtraPackageConnectionEN     ErrorType = "EREP0000 " + failureConnectionEN + " (get extra package)"
	ExtraPackageConnectionID     ErrorType = "EREP0000 " + failureConnectionID + " (get extra package)"
	ExtraPackageParsingFailedEN  ErrorType = "EREP0001 " + parsingFailedEN + " (get extra package)"
	ExtraPackageParsingFailedID  ErrorType = "EREP0001 " + parsingFailedID + " (get extra package)"
	ExtraPackageFeeResponseNotOK ErrorType = "EREP0002 " + ResponseNotOk + " %s"
	ExtraPackageError            ErrorType = "EREP0003 %s"
	ExtraPackageInternalError    ErrorType = "EREP0004 " + internalError

	// ERROR PROMO
	PromoError                         ErrorType = "ERPR0000 %s"
	PromoValidationError               ErrorType = "ERPR0001 %s"
	PromoPGWError                      ErrorType = "ERPR0002 %s"
	PromoPGWParsingError               ErrorType = "ERPR0003 %s"
	PromoPGWResponseNotOk              ErrorType = "ERPR0004 %s"
	PromoPGWConnection                 ErrorType = "ERPR0005 %s"
	PromoPGWPAutoApplyParisngError     ErrorType = "ERPR0006 %s"
	PromoPGWRAutoApplyesponseNotOk     ErrorType = "ERPR0007 %s"
	PromoShipmentSettingError          ErrorType = "ERPR0008 %s"
	PromoAutoApplyShipmentSettingError ErrorType = "ERPR0009 %s"
	PromoAutoApplyError                ErrorType = "ERPR0010 %s"
	PromoValidationApiError            ErrorType = "ERPR0011 %s"
	PromoValidationInputFormat         ErrorType = "ERPR0012 %s"
	PromoValidationResponseFormat      ErrorType = "ERPR0013 %s"
	PromoValidationFormat              ErrorType = "ERPR0014 %s"
	PromoValidationSubmit              ErrorType = "ERPR0015 %s" //WILL BE SHOWN 403 ON ERROR_CODE

	// GET NEAREST WASTE BANK (PAXEL RECYCLE) MAIN
	GNWBMFailureConnectionEN ErrorType = "ERGNWBM0000 " + failureConnectionEN + " (get nearest waste banks main)"
	GNWBMFailureConnectionID ErrorType = "ERGNWBM0000 " + failureConnectionID + " (get nearest waste banks main)"
	GNWBMInternalError       ErrorType = "ERGNWBM0001 " + internalError
	GNWBMError               ErrorType = "ERGNWBM0002 %s"
	GNWBNotCoverageError     ErrorType = "ERGNWBM0003 Delivery address for Paxel Recycle not coverage"

	// ALLOW PAXEL RECYCLE by PICKUP DATE MAIN
	APRMFailureConnectionEN ErrorType = "ERAPRM0000 " + failureConnectionEN + " (allow paxel recycle by pickup date main)"
	APRMFailureConnectionID ErrorType = "ERAPRM0000 " + failureConnectionID + " (allow paxel recycle by pickup date main)"
	APRMInternalError       ErrorType = "ERAPRM0001 " + internalError
	APRMError               ErrorType = "ERAPRM0002 %s"
	APRMMemberCodeMandatory ErrorType = "ERAPRM0003 Member code must to be filled"

	// PAYLOAD FIELDS
	ServiceTypeError ErrorType = "ERPL0001 Invalid input service type"
	EstimatedDateID  ErrorType = "ERPL0002 Tanggal estimasi %s salah"
	EstimatedDateEN  ErrorType = "ERPL0003 Wrong estimated date %s"
)
