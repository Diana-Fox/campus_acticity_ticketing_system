package main

import (
	"log"

	"campus_acticity_ticketing_system/ticket_admin/internal/app"
)

func main() {
	// 初始化配置
	if err := initConfig(); err != nil {
		log.Fatalf("初始化配置失败: %s", err)
	}

	// 启动应用
	if err := app.Run(); err != nil {
		log.Fatalf("应用启动失败: %s", err)
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	return viper.ReadInConfig()
}
