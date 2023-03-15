package main

import (
	"fmt"
	"log"

	todo "github.com/B-danik/GameShop_Go"
	"github.com/B-danik/GameShop_Go/cmd/api"
	"github.com/B-danik/GameShop_Go/cmd/appDB"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Run Test Application: GameShop_Backend...")

	router, err := api.Header()
	if err != nil {
		log.Fatalf("Error Header... %s", err.Error())
	}
	appDB.Connect()

	log.Println("Start app")
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("Port"), router); err != nil {
		log.Fatalf("Error conect")
	}
}
