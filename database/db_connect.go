package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"task1/model"
)


var db *gorm.DB

func Setup() *gorm.DB {
	var err error
	db, err = gorm.Open("mysql", "root:anitha@tcp(localhost:3306)/employee_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Address{}, &models.Contact{}, &models.Employee{})
	return db
}
