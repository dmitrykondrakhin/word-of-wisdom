package config

import "github.com/spf13/viper"

type Config struct {
	ServerHost    string `mapstructure:"SERVER_HOST"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	ClientHost    string `mapstructure:"CLIENT_HOST"`
	ClientPort    string `mapstructure:"CLIENT_PORT"`
	RepeatedCount int    `mapstructure:"REPEATED_COUNT"`
	HashCashBits  uint   `mapstructure:"HASHCASH_BITS"`
}

func Init() (config Config, err error) {
	viper.SetDefault("SERVER_HOST", "127.0.0.1")
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("CLIENT_HOST", "127.0.0.1")
	viper.SetDefault("CLIENT_PORT", "8080")
	viper.SetDefault("REPEATED_COUNT", "60")
	viper.SetDefault("HASHCASH_BITS", "20")

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}
