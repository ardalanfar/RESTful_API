package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname_u string `json:"fullname_u,omitempty"`
	Age_u      int    `json:"age_u,omitempty"`
	Courseid   int    `json:"courseid,omitempty"`
}

func UserMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
