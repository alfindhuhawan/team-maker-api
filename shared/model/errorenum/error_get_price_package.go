package errorenum

const (
	// ERROR GET PRICE PACKAGE
	// FindMemberPGWFailureConnectionEN ErrorType = "ERFMP0000 " + failureConnectionEN
	// FindMemberPGWParsingFailedEN     ErrorType = "ERFMP0001 " + parsingFailedEN + " (find member pgw)"
	// FindMemberPGWParsingFailedID     ErrorType = "ERFMP0001 " + parsingFailedID + " (find member pgw)"
	// FindMemberPGWResponseNotOK       ErrorType = "ERFMP0002 " + ResponseNotOk + " %s"
	// FindMemberPGWInternalError       ErrorType = "ERFMP0003 %s"
	GetPricePackageSizeNotDefinedEN ErrorType = "ERGPP0001 Size Not Defined"
	GetPricePackageSizeNotDefinedID ErrorType = "ERGPP0001 Size Tidak Terdefinisi"
)
