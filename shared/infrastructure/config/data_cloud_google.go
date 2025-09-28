package config

type CloudGoogle struct {
	PrivateKey     string `json:"private_key"`
	ClientEmail    string `json:"client_email"`
	ExpireInMinute int64  `json:"expire_in_minute"`
}
