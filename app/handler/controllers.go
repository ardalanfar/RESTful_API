package handler

import (
	"Service_Restful/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Show_usercourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var users []models.User
	db.Model(&models.User{}).Select("users.fullname_u,users.age_u").Joins("inner join courses on courses.id = users.courseid").Where("courses.name_c = ?", params["name_c"]).Find(&users)
	json.NewEncoder(w).Encode(users)
}
