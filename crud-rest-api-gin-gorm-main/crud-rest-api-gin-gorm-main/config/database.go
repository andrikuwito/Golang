package config

import (
	"github.com/jinzhu/gorm"
	"rest-api-km/structs"
)

func InitDB() *gorm.DB {
	connection := "root:root@tcp(127.0.0.1:3306)/hacktiv_restapi?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		panic("Failed connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}