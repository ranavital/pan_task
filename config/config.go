package config

var AppConfig = &Config{}

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
}