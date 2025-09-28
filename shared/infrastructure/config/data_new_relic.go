package config

type NewRelic struct {
	Enable                         bool   `json:"enable"`
	EnableIncomingLogging          bool   `json:"enable_incoming_logging"`
	EnableOutgoingLogging          bool   `json:"enable_outgoing_logging"`
	EnableLoggingOnSuccessIncoming bool   `json:"enable_logging_on_success_incoming"`
	EnableLoggingOnSuccessOutgoing bool   `json:"enable_logging_on_success_outgoing"`
	NameBuyer                      string `json:"name_buyer"`
	NameSeller                     string `json:"name_seller"`
	NameCore                       string `json:"name_core"`
	NameCMS                        string `json:"name_cms"`
	Key                            string `json:"key"`
}
