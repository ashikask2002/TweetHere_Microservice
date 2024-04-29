package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	AdminSvcUrl string `mapstructure:"ADMIN_SVC_URL"`
}

var envs = []string{
	"PORT", "ADMIN_SVC_URL",
}

// func LoadConfig() (Config, error) {
// 	fmt.Println("inside LoadConfig")
// 	var config Config

// 	viper.AddConfigPath("/home/ashik/Git/TweetHere_Microservice/TweetHere-API-GATEWAY/")
// 	viper.SetConfigName(".env")

// 	if err := viper.ReadInConfig(); err != nil {
// 		return config, fmt.Errorf("failed to read config file: %v", err)
// 	}

// 	for _, env := range envs {
// 		if err := viper.BindEnv(env); err != nil {
// 			return config, fmt.Errorf("failed to bind environment variable %s: %v", env, err)
// 		}
// 	}

// 	if err := viper.Unmarshal(&config); err != nil {
// 		return config, fmt.Errorf("failed to unmarshal config: %v", err)
// 	}

// 	fmt.Println("config loaded successfully:", config)
// 	return config, nil
// }

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
	return config, nil
}
