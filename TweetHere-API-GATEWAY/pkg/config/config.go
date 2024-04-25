package config

import "github.com/spf13/viper"


type Config struct{
	Port    string `mapstructure:"PORT"`
	AdminSvcUrl string `mapstructure:"ADMIN_SVC_URL"`
}

var envs = []string{
	"PORT", "ADMIN_SVC_URL",
}

func LoadConfig() (Config,error){
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.ReadInConfig()


	for _,env := range envs{
		if err := viper.BindEnv(env); err != nil{
			return config,err
		}
	}

	if err := viper.Unmarshal(&config);err != nil{
		return config,err
	}
	return config,nil
}