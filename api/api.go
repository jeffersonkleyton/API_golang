package api

import (
	"fmt"
	"log"
	"net/http"
	"testando/api/models"
	"testando/api/routes"
)

func Run() {
	db := models.Connect()
	defer db.Close()

	if !db.HasTable(&models.User{}) {
		db.Debug().Create(&models.User{})
	}
	listen(3000)
}

func listen(p int) {
	fmt.Printf(" executando na porta: 3000")
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
