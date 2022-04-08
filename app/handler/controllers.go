package handler

import (
	"RESTful_API/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Report struct {
	Student string
	Course  string
	Teacher string
}

func Get_usercourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var users []models.Users
	db.Model(&models.Users{}).Select("users.fullname_u,users.age_u").Joins("inner join courses on courses.id = users.courseid").Where("courses.name_c = ?", param["name_c"]).Find(&users)
	//select users.fullname_u,users.age_u from users inner join courses on users.courseid = courses.id where courses.name_c = ["name_c"];
	json.NewEncoder(w).Encode(users)
}

func Get_report(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var report []Report
	db.Model(&models.Users{}).Select("users.fullname_u as Student ,courses.name_c as Course ,teachers.fullname_t as Teacher").Joins("inner join courses on courses.id = users.courseid").Joins("inner join teachers on teachers.id = courses.teacher_id").Find(&report)
	//select users.fullname_u as Student ,courses.name_c as Courses ,teachers.fullname_t as Teacher
	//from users
	//inner join courses
	//on courses.id = users.courseid
	//inner join teachers
	//on teachers.id = courses.teacher_id
	json.NewEncoder(w).Encode(report)
}
