package config

type AuthKey struct {
	KeyConfig string `json:"key_config"`
	JWTSecret string `json:"jwt_secret"`
}
