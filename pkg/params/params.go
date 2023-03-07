package params

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitParams() {
	viper.SetDefault("http_port", ":8000")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Using defualt config...")
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

	}
}
