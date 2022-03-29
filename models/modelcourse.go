package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name_c    string `json:"name_c,omitempty"`
	TeacherID int    `json:"teacherid,omitempty"`
}

func CourseMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Course{})
	return db
}
