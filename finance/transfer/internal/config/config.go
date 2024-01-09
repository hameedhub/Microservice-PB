package config

import "github.com/spf13/viper"

type Config struct {
	SERVER_IDLE_TIMEOUT       int64  `yaml:"SERVER_IDLE_TIMEOUT"`
	SERVER_READ_TIMEOUT       int64  `yaml:"SERVER_READ_TIMEOUT"`
	SERVER_WRITE_TIMEOUT      int64  `yaml:"SERVER_WRITE_TIMEOUT"`
	SERVER_PORT               string `yaml:"SERVER_PORT"`
	KAFKA_SERVER              string `yaml:"KAFKA_SERVER"`
	KAFKA_CONSUMER_GROUP      string `yaml:"KAFKA_CONSUMER_GROUP"`
	HIGH_PRIORITY_PARTITION   int64  `yaml:"HIGH_PRIORITY"`
	MEDIUM_PRIORITY_PARTITION int64  `yaml:"MEDIUM_PRIORITY"`
	LOW_PRIORITY_PARTITION    int64  `yaml:"LOW_PRIORITY"`
}

// https://github.com/spf13/viper
func ReadConfig(path string) (c Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)

	return
}
