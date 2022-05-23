package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	URL := "root:root@tcp(127.0.0.1:3306)/go_teste"
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
