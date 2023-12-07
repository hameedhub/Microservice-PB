package config

import "github.com/spf13/viper"

type Config struct {
	DB_SOURCE            string `mapstructure:"DB_SOURCE"`
	SERVER_IDLE_TIMEOUT  int64  `mapstructure:"SERVER_IDLE_TIMEOUT"`
	SERVER_READ_TIMEOUT  int64  `mapstructure:"SERVER_READ_TIMEOUT"`
	SERVER_WRITE_TIMEOUT int64  `mapstructure:"SERVER_WRITE_TIMEOUT"`
	SERVER_PORT          int64  `mapstructure:"SERVER_PORT"`
	KAFKA_SERVER         string `mapstructure:"KAFKA_SERVER"`
	KAFKA_PRODUCER_TOPIC string `mapstructure:"KAFKA_PRODUCER_TOPIC"`
	KAFKA_CONSUMER_TOPIC string `mapstructure:"KAFKA_CONSUMER_TOPIC"`
	KAFKA_CONSUMER_GROUP string `mapstructure:"KAFKA_CONSUMER_GROUP"`
}

// https://github.com/spf13/viper
func ReadConfig(path string) (c Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)

	return
}
