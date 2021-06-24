package data

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DatabaseConection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loding .env file!")
	}

	dsn := os.Getenv("DB_CONNECT")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB接続完了")
}
