package configs

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

type envConfigs struct {
	SECRET_KEY                 string
	ACCESS_TOKEN_EXPIRE_MINUES int

	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
	DATABASE_SSL_MODE string
}

func loadEnvVariables() *envConfigs {

	var config envConfigs
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return &config

}
