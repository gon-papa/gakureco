package data

import (
	"crypto/sha256"
	"fmt"
	"io"
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

// ストレッチングしていない
func Encrypt(plaintext string) (crypttext string) {
	cryptext := sha256.New()
	io.WriteString(cryptext, plaintext)
	crypttext = fmt.Sprintf("%x", cryptext.Sum(nil))
	return
}
