package init

import "github.com/spf13/viper"

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	return viper.ReadInConfig()
}
