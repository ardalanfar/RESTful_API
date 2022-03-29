package database

import (
	"Service_Restful/conf"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	config := conf.GetConfig()

	dsn := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Dbname,
		config.DB.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect database")
	}
	return db
}
