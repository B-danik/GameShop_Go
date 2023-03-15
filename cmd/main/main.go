package main

import (
	"fmt"
	"log"
	"os"

	todo "github.com/B-danik/GameShop_Go"
	"github.com/B-danik/GameShop_Go/internal/database"
	"github.com/B-danik/GameShop_Go/internal/user"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Run Test Application: GameShop_Backend...")

	log.Println("Create router")
	router := httprouter.New()

	log.Println("Create handler")
	handler := user.NewHandler()
	handler.Register(router)

	if err := initConfig(); err != nil {
		log.Fatalf("Error initConfigYaml... %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error initConfigEnv... %s", err.Error())
	}

	fmt.Println(viper.GetString("db.port"))
	db, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("db_password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Error connect db... %s", err.Error())
	}

	log.Printf("DB name: %s", db.DriverName())
	log.Println("Start app")
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("Port"), router); err != nil {
		log.Fatalf("Error conect")
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
