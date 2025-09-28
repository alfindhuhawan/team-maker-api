package enum

type PaymentMethod string

const (
	PaxelCreditEnum        = PaymentMethod("CRD")
	PaxelCashOnPickoupEnum = PaymentMethod("COP")
)
