package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Fullname_u string `json:"fullname_u,omitempty"`
	Age_u      int    `json:"age_u,omitempty"`
	Courseid   int    `json:"courseid,omitempty"`
}

func UsersMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Users{})
	return db
}
