package vo

type BankInfo struct {
	BankCode        string `json:"bank_code" bson:"bank_code"`
	BankName        string `json:"bank_name" bson:"bank_name"`
	AccountName     string `json:"account_name" bson:"account_name"`
	AccountNumber   string `json:"account_number" bson:"account_number"`
	AccountVerified bool   `json:"account_verified" bson:"account_verified"`
}
