package config

import "github.com/spf13/viper"

func InitConfig(file string) error {

	viper.SetConfigFile(file)

	err := viper.ReadInConfig()
	return err
}
