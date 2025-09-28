package errorenum

const (
	// ERROR FIND MEMBER PGW
	FindMemberPGWFailureConnectionEN ErrorType = "ERFMP0000 " + failureConnectionEN + " (find member pgw)"
	FindMemberPGWFailureConnectionID ErrorType = "ERFMP0000 " + failureConnectionID + " (find member pgw)"
	FindMemberPGWParsingFailedEN     ErrorType = "ERFMP0001 " + parsingFailedEN + " (find member pgw)"
	FindMemberPGWParsingFailedID     ErrorType = "ERFMP0001 " + parsingFailedID + " (find member pgw)"
	FindMemberPGWResponseNotOK       ErrorType = "ERFMP0002 " + ResponseNotOk + " %s"
	FindMemberPGWError               ErrorType = "ERFMP0003 %s"
	FindMemberPGWInternalError       ErrorType = "ERFMP0004 " + internalError
)
