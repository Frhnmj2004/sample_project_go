package config

import "github.com/spf13/viper"

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func GetConfig() *viper.Viper {
	return viper.GetViper()
}
