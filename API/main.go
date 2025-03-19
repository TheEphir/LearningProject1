package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func make_connection() {
	connStr := "host=192.168.0.165 user=theephir password=megamax05 dbname=dvdrental port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	var version string

	err = db.Raw("SELECT version()").Scan(&version).Error
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(version)
}

func main() {
	make_connection()
}
