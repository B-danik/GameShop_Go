package main

import (
	"fmt"
	"log"

	todo "github.com/B-danik/GameShop_Go"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello world")
	if err := initConfig(); err != nil {
		log.Fatalf("Error initConfig %s", err.Error())
	}
	fmt.Println(viper.GetString("Port"))
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("Port")); err != nil {
		log.Fatalf("Error conect")
	}
}

func initConfig() error {
	viper.AddConfigPath("/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
