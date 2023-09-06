package config

import "github.com/spf13/viper"

type Config struct {
	JWTKey   string
	DBUrl    string
	DBDriver string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return &Config{
		JWTKey:   viper.GetString("JWTKEY"),
		DBUrl:    viper.GetString("DB_URL"),
		DBDriver: viper.GetString("DB_DRIVER"),
	}
}
