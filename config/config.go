package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	JWTKey   string `mapstructure:"JTWKEY"`
	DBUrl    string `mapstructure:"DB_URL"`
	DBDriver string `mapstructure:"DB_DRIVER"`
}

var C Config

func LoadConfig(filePath string) {
	viper.SetConfigFile(filePath)
	viper.ReadInConfig()

	if err := viper.Unmarshal(&C); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
