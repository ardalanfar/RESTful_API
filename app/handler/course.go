package handler

import (
	"Service_Restful/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Getcourses(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var courses []models.Course
	db.Find(&courses)
	json.NewEncoder(w).Encode(courses)
}

func Getcourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var course models.Course
	db.First(&course, params["id"])
	json.NewEncoder(w).Encode(course)
}

func Createcourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course models.Course
	json.NewDecoder(r.Body).Decode(&course)
	db.Create(&course)
	json.NewEncoder(w).Encode(course)
}
