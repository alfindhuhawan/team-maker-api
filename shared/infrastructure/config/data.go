package config

type Config struct {
	Token                string              `json:"token"`
	JWTSecretKey         string              `json:"jwt_secret_key"`
	Database             Database            `json:"database"`
	Cache                Cache               `json:"cache"`
	MessageBroker        MessageBroker       `json:"message_broker"`
	ApplicationServer    ApplicationServer   `json:"application_server"`
	ApplicationExternal  ApplicationExternal `json:"application_external"`
	AuthKey              AuthKey             `json:"auth_key"`
	NewRelic             NewRelic            `json:"new_relic"`
	ExceptionServiceType []string            `json:"exception_service_type"`
	GCS                  GCS                 `json:"gcs"`
	CloudGoogle          CloudGoogle         `json:"cloud_google"`
}
