package configs

import (
	"log"

	"github.com/spf13/viper"
)

type EnvCOnfig struct {
	Port string `mapstructure:"PORT"`
}

func Loadenv() (config *EnvCOnfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error readinmg env file", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
