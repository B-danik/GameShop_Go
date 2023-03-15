package appDB

import (
	"fmt"
	"log"
	"os"

	"github.com/B-danik/GameShop_Go/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type appDB struct {
}

func Connect() error {
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
	return nil
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
