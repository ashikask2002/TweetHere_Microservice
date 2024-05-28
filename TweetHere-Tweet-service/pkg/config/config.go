package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost                string `mapstructure:"DB_HOST"`
	DBname                string `mapstructure:"DB_NAME"`
	DBUser                string `mapstructure:"DB_USER"`
	DBPort                string `mapstructure:"DB_PORT"`
	DBPassword            string `mapstructure:"DB_PASSWORD"`
	Port                  string `mapstructure:"PORT"`
	AUTH_SVC_URL          string `mapstructure:"AUTH_SVC_URL"`
	REGION                string `mapstructure:"REGION"`
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	KafkaPort             string `mapstructure:"KAFKA_PORT"`
	KafkaTpic             string `mapstructure:"KAFKA_TOPIC"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "PORT", "AUTH_SVC_URL", "KAFKA_PORT", "KAFKA_TOPIC",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	fmt.Println("configggg", config)
	return config, nil
}
