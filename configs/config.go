package configs

import (
	"log"

	pg "github.com/Heian28/go-utils/db/gopostgres"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
	Port int    `mapstructure:"port"`
}

type Config struct {
	App      ServiceConfig              `mapstructure:"app"`
	Postgres pg.GoPostgresConfiguration `mapstructure:"postgres"`
}

var AppConfig Config

func loadConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}
	return nil
}

func Init() {
	if err := loadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}
