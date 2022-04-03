package config

var AppConfig = &Config{}

type Config struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	DBPath string `json:"db_path"`
}
