package configs

var cfg *config

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Database string
}

type config struct {
	API APIConfig
	DB DBConfig
}