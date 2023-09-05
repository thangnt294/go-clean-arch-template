package config

import "github.com/spf13/viper"

type Config struct {
	JWTKey string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return &Config{
		JWTKey: viper.GetString("JWTKEY"),
	}
}
