package util

import "github.com/spf13/viper"

type Config struct {
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
	AdiqKey     string `mapstructure:"ADIQ_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
