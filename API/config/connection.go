package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func getDbCreds() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Can't load .env file")
	}
	return map[string]string{
		"dbHost":   os.Getenv("DB_HOST"),
		"dbUser":   os.Getenv("DB_USER"),
		"dbPasswd": os.Getenv("DB_PASSWORD"),
		"dbName":   os.Getenv("DB_NAME"),
		"dbPort":   os.Getenv("PORT"),
	}
}

func Connect() {
	dbCreds := getDbCreds()
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbCreds["dbHost"], dbCreds["dbUser"], dbCreds["dbPasswd"], dbCreds["dbName"], dbCreds["dbPort"],
	)
	d, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = d // 13 row
}

func GetDb() *gorm.DB {
	return db
}
