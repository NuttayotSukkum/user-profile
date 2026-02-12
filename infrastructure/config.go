package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Configs

type Configs struct {
	App AppConfigs
}

type AppConfigs struct {
	Port       string
	HmacSecret []byte
}

func InitConfig() {
	var cfg Configs
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	AppConfig = &cfg

	log.Println("Config loaded")
}
