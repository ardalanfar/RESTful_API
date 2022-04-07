package handler

import (
	"RESTful_API/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Getusers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.Users
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func Getuser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var user models.Users
	db.First(&user, param["id"])
	json.NewEncoder(w).Encode(user)
}

func Createuser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}
