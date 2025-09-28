package config

type RabbitMQ struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type MessageBroker struct {
	Rabbitmq RabbitMQ `json:"rabbitmq"`
}
