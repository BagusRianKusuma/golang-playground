package config

type Config struct {
	WeatherAPIKey string
}

func LoadConfig() *Config {
	return &Config{
		WeatherAPIKey: "4903e14e0e084a1c855114812251303",
	}
}
