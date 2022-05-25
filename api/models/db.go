package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func Connect() *gorm.DB {
	godotenv.Load()
	port := os.Getenv("HOST_DB")
	URL := "root:root@tcp(" + port + ":3306)/go_teste"
	fmt.Println("URL", URL)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
