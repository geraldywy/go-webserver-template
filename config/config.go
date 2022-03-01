package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	 DBHost string `mapstructure:"db_host"`
	 DBUser string `mapstructure:"db_user"`
	 DBPass string `mapstructure:"db_pass"`
	 DBName string `mapstructure:"db_name"`
}

var Conf Config

func InitConfig(configDir, configName, configExt string) {
	fmt.Println("initializing config")
	viper.AddConfigPath(configDir)
	viper.SetConfigName(configName)
	viper.SetConfigType(configExt)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file not found")
		} else {
			// Config file was found but another error was produced
			panic(fmt.Sprintf("unexpected err, reading config file: %v", err))
		}
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
}