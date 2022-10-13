package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func init() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Failed to read env variable")
		fmt.Print(errEnv)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", dbHost, dbUserName, dbPassword, dbName, dbPort)
	fmt.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
		fmt.Print(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
