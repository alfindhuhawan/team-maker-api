package enum

type ProgressFirstmileOrLastMileEnum string

const (
	SuccessEarlyProgressFirstmileOrLastMileEnum    = ProgressFirstmileOrLastMileEnum("Success-Early")
	SuccessOntimeProgressFirstmileOrLastMileEnum   = ProgressFirstmileOrLastMileEnum("Success-OnTime")
	SuccessLateProgressFirstmileOrLastMileEnum     = ProgressFirstmileOrLastMileEnum("Success-Late")
	RejectProgressFirstmileOrLastMileEnum          = ProgressFirstmileOrLastMileEnum("Reject")
	CancelledProgressFirstmileOrLastMileEnum       = ProgressFirstmileOrLastMileEnum("Cancelled")
	OnprocessOntimeProgressFirstmileOrLastMileEnum = ProgressFirstmileOrLastMileEnum("OnProcess-Ontime")
	OnprocessLateProgressFirstmileOrLastMileEnum   = ProgressFirstmileOrLastMileEnum("OnProcess-Late")
	UndeliveryProgressFirstmileOrLastMileEnum      = ProgressFirstmileOrLastMileEnum("Undelivery")
	FailedProgressFirstmileOrLastMileEnum          = ProgressFirstmileOrLastMileEnum("Failed")
)
