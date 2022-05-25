package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testando/api/models"
	"testando/api/utils"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id    uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name  string `gorm:"size:100; not null" json:"name"`
	Email string `gorm:"size:100; not null" json:"email"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetUser()
	utils.Response(w, users, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	user := models.GetUserById(id)
	utils.Response(w, user, http.StatusOK)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	var err error
	body := utils.BodyParser(r)
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	userDB, err := models.GetUserByEmail(user.Email)
	if userDB.Email == user.Email {
		utils.Response(w, "Informe um email válido", http.StatusCreated)
		return
	}
	if err != nil {
		utils.Response(w, err.Error(), http.StatusUnprocessableEntity)
	}
	sb, err := bcrypt.GenerateFromPassword([]byte(string(user.Password)), 10)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user.Password = string(sb)
	err = models.NewUser(user)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.Response(w, "Novo usuário criado", http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user.Id = id
	rows, err := models.UpdateUserById(user)

	if err != nil {
		utils.Response(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.Response(w, rows, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	_, err := models.DeleteUser(id)
	if err != nil {
		utils.Response(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
