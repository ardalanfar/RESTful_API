package handler

import (
	"RESTful_API/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Get_usercourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var users []models.Users
	db.Model(&models.Users{}).Select("users.fullname_u,users.age_u").Joins("inner join courses on courses.id = users.courseid").Where("courses.name_c = ?", param["name_c"]).Find(&users)
	//select users.fullname_u,users.age_u from users inner join courses on users.courseid = courses.id where courses.name_c = ["name_c"];
	json.NewEncoder(w).Encode(users)
}

func Get_users_inTeacher(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//param := mux.Vars(r)
	var users []models.Users
	//db.Model(&models.Users{}).Select("users.fullname_u,users.age_u").Joins("inner join courses on courses.id = users.courseid").Where("courses.name_c = ?", param["name_t"]).Find(&users)
	json.NewEncoder(w).Encode(users)
}

//select users.fullname_u as Student ,courses.name_c as Courses ,teachers.fullname_t as Teacher
//from ((users
//inner join courses on users.courseid = courses.id)
//inner join teachers on courses.teacher_id = teachers.id);
