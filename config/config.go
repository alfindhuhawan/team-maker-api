package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	MongoURI  string `json:"mongo_uri"`
	MongoDB   string `json:"mongo_db"`
	JwtSecret string `json:"jwt_secret"`
}

var AppConfig *Config

func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := &Config{}
	if err := decoder.Decode(cfg); err != nil {
		log.Fatalf("Cannot decode config: %v", err)
	}
	AppConfig = cfg
	log.Println("Config loaded")
}
