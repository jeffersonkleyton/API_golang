package api

import (
	"log"
	"net/http"
	"testando/api/models"
	"testando/api/routes"
)

func Run() {
	db := models.Connect()
	defer db.Close()

	if !db.HasTable(&models.User{}) {
		db.Debug().CreateTable(&models.User{})
	}
	listen(3000)
}

func listen(p int) {
	//port := fmt.Sprintf(":%d", p)
	//fmt.Printf("\n\nListening port %s...\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}
