package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Fullname_t string `json:"fullname_t,omitempty"`
	Email_t    string `json:"email_t,omitempty"`
}

func TeacherMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Teacher{})
	return db
}
